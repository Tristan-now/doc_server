package image_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
	"gvd_server/service/common/list"
	"gvd_server/service/common/res"
)

type ImageListResponse struct {
	models.ImageModel
	WebPath string `json:"webPath"`
	//NickName string `json:"nickName"`
}

func (ImageApi) ImageListView(c *gin.Context) {
	var cr models.Pagination
	c.ShouldBindJSON(&cr)
	_list, count, _ := list.QueryList(models.ImageModel{}, list.Option{
		Pagination: cr,
		Likes:      []string{"fileName"},
		//Preload:    []string{"UserModel"},
	})
	var imageList = make([]ImageListResponse, 0)
	for _, model := range _list {
		imageList = append(imageList, ImageListResponse{
			model,
			model.WebPath(),
			//model.UserModel.NickName,
		})
	}
	res.OKWithList(imageList, count, c)
}
