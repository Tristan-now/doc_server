package image_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"path"
	"strings"
)

var ImageWhiteList = []string{
	"jpg",
	"png",
	"jpeg",
	"gif",
	"svg",
	"webp",
}

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

	if !InImageWhiteList(file.Filename, ImageWhiteList) {
		res.FailWithMsg("文件非法", c)
		//log.Warn("文件非法")
		return
	}

	path := path.Join("uploads", claims.NickName, file.Filename)
	c.SaveUploadedFile(file, path)
	res.OKWithData("图片上传成功", c)

}

func InImageWhiteList(fileName string, whiteList []string) bool {
	// 截取文件后缀
	_list := strings.Split(fileName, ".") // xxx  1.2 xxx.png xxx.PNG  xxx.png   xxx.1.2.png  xxxx.png.exe
	if len(_list) < 2 {
		return false
	}
	suffix := strings.ToLower(_list[len(_list)-1])
	for _, s := range whiteList {
		if suffix == s {
			return true
		}
	}
	return false
}
