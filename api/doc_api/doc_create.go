package doc_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/full_search_service"
	"strings"
)

type DocCreateRequest struct {
	Title    string `json:"title" binding:"required" label:"文章标题"`
	Content  string `json:"content" binding:"required" label:"文章内容"`
	ParentID *uint  `json:"parentID"`
}

// DocCreateView 创建文档
// @Tags 文档管理
// @Summary 创建文档
// @Description 创建文档，创建成功之后，data=文档id
// @Param data body DocCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/docs [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (DocApi) DocCreateView(c *gin.Context) {
	var cr DocCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// 判断ParentID是不是合法的
	if cr.ParentID != nil {
		if *cr.ParentID <= 0 {
			res.FailWithMsg("父文档ID非法", c)
			return
		}
		var parentModel models.DocModel
		err = global.DB.Take(&parentModel, *cr.ParentID).Error
		if err != nil {
			res.FailWithMsg("父文档不存在", c)
			return
		}
	}

	var docModel = models.DocModel{
		Title:    cr.Title,
		Content:  cr.Content,
		ParentID: cr.ParentID,
	}

	err = global.DB.Create(&docModel).Error
	if err != nil {
		res.FailWithMsg("文档保存失败", c)
		return
	}
	go full_search_service.FullSearchCreate(docModel)

	var docList []models.DocModel
	models.FindAllParentDocList(docModel, &docList)
	var docIDList []string
	docLen := len(docList)
	for i := docLen - 1; i >= 0; i-- {
		docIDList = append(docIDList, fmt.Sprintf("%d", docList[i].ID))
	}
	key := strings.Join(docIDList, ".")
	fmt.Println("计算的key = " + key)
	global.DB.Model(&docModel).Update("key", key)
	res.OK(docModel.ID, "文档添加成功", c)
}
