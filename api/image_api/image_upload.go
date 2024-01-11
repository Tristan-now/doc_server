package image_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"path"
)

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Param token header string true "token"
// @Param image formData file true "文件上传"
// @Router /api/image [post]
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageApi) ImageUploadView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		res.FailWithMsg("图片参数错误", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	path := path.Join("uploads", claims.NickName, file.Filename)
	c.SaveUploadedFile(file, path)
	res.OKWithData("图片上传成功", c)

}
