package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"
	"strings"
)

// DocContentView 文档内容
// @Tags 文档管理
// @Summary 文档内容
// @Description 文档内容
// @Param id path int true "id"
// @Router /api/docs/{id} [get]
// @Produce json
// Success 200 {object} res.Response{data=DocContentResponse}
func (DocApi) DocContentView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}

	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	var roleID uint = 2
	var userID uint = 0

	if err == nil {
		roleID = claims.RoleID
		userID = claims.UserID
	} else {
		res.FailWithError(err, c)
		return
	}

	docContent, ok := redis_service.GetDocContent(roleID, cr.ID, userID)
	if ok {
		res.OKWithData(docContent, c)
		return
	}

	//判断角色是否拥有文档的访问权限
	var roleDoc models.RoleDocModel
	err = global.DB.
		Preload("DocModel.UserCollDocList").
		Preload("RoleModel").
		Take(&roleDoc, "role_id = ? and doc_id = ?", roleID, cr.ID).Error

	if err != nil {
		res.FailWithMsg("文档鉴权失败", c)
		return
	}

	redis_service.NewDocLook().SetById(cr.ID)

	doc := roleDoc.DocModel
	docDigg := redis_service.NewDocDigg().GetById(doc.ID)
	docLook := redis_service.NewDocLook().GetById(doc.ID)

	var response = redis_service.DocContentResponse{
		DiggCount: docDigg + doc.DiggCount,
		LookCount: docLook + doc.LookCount,
		CollCount: len(doc.UserCollDocList),
		Title:     doc.Title,
		CreatedAt: doc.CreatedAt,
	}

	isDocFree := strings.Contains(doc.Content, global.DocSplitSign)

	var freeContent string                                                 //试看部分
	var content = strings.ReplaceAll(doc.Content, global.DocSplitSign, "") //实际正文
	if isDocFree {
		_list := strings.Split(doc.Content, global.DocSplitSign)
		freeContent = _list[0]
	}

	if roleDoc.FreeContent != nil {
		response.IsSee = true

		if doc.FreeContent != "" {
			freeContent = doc.FreeContent
		}

		if *roleDoc.FreeContent != "" {
			freeContent = *roleDoc.FreeContent
		}
	}

	// IsPwd  判断这个角色有没有密码
	if roleDoc.Pwd != nil && (*roleDoc.Pwd != "" || roleDoc.RoleModel.Pwd != "") {
		response.IsPwd = true
	}

	// IsColl
	if roleID != 2 {
		// 查用户是否收藏了文档
		var userDoc models.UserCollDocModel
		err = global.DB.Take(&userDoc, "doc_id = ? and user_id = ?", cr.ID, claims.UserID).Error
		if err == nil {
			response.IsColl = true
		}
		// 用户是否对这个文档免密
		var usePwd models.UserPwdDocModel
		err = global.DB.Take(&usePwd, "doc_id = ? and user_id = ?", cr.ID, claims.UserID).Error
		if err == nil {
			response.IsPwd = false
		}
	}

	// Content
	// 有密码。有试看  试看内容
	// 无密码，有试看  试看内容
	if response.IsSee {
		response.Content = freeContent
	}

	// 有密码，无试看  空
	if response.IsPwd && !response.IsSee {
		response.Content = ""
	}

	// 无密码，无试看  正文
	if !response.IsPwd && !response.IsSee {
		response.Content = content
	}

	redis_service.SetDocContent(roleID, cr.ID, userID, response)

	res.OKWithData(response, c)
}
