package image_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/hash"
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
	fileHeader, err := c.FormFile("image")
	if err != nil {
		res.FailWithMsg("图片参数错误", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	savePath := path.Join("uploads", claims.NickName, fileHeader.Filename)
	//文件名后缀白名单校验
	if !InImageWhiteList(fileHeader.Filename, ImageWhiteList) {
		res.FailWithMsg("文件非法", c)
		//log.Warn("文件非法")
		return
	}

	//文件大小校验
	if fileHeader.Size > int64(4*1024*1024) {
		res.FailWithMsg("文件过大", c)
		return
	}

	//计算文件hash
	file, _ := fileHeader.Open()
	fileHash := hash.FileMd5(file)

	//对重复文件的判断
	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, "hash = ?", fileHash).Error
	// 没有 要上传，要入库
	// 有 只需要入库，但是入库的path需要改成和有的那个一样
	if err != nil {
		err = c.SaveUploadedFile(fileHeader, savePath)
		if err != nil {
			global.Log.Errorf("%s 文件保存错误%s", savePath, err)
			res.FailWithMsg("上传图片错误", c)
			return
		}
	} else {
		savePath = imageModel.Path

	}

	imageModel = models.ImageModel{
		UserID:   claims.UserID,
		FileName: fileHeader.Filename,
		Size:     fileHeader.Size,
		Path:     savePath,
		Hash:     fileHash,
	}

	err = global.DB.Create(&imageModel).Error
	if err != nil {
		global.Log.Errorln(err)
		//log.Error("图片上传失败")
		res.FailWithMsg("文件上传失败", c)
		return
	}
	err = c.SaveUploadedFile(fileHeader, savePath)
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
