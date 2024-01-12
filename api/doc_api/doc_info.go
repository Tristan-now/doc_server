package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

type DocInfoResponse struct {
	Model         models.Model
	Title         string `json:"title"`
	ContentLength int    `json:"contentLength"`
	DiggCount     int    `json:"diggCount"`
	LookCount     int    `json:"lookCount"`
	Key           string `json:"key"`
}

func (DocApi) DocInfoView(c *gin.Context) {
	var cr models.IDRequest
	c.ShouldBindUri(&cr)

	var doc models.DocModel
	err := global.DB.Take(&doc, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}

	var docInfo = DocInfoResponse{
		Model:         doc.Model,
		Title:         doc.Title,
		ContentLength: len(doc.Content),
		//DiggCount:     docDigg + doc.DiggCount,
		//LookCount:     docLook + doc.LookCount,
		Key: doc.Key,
	}
	res.OKWithData(docInfo, c)
}
