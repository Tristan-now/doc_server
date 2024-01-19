package data_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"time"
)

type DataCountResponse struct {
	DateList  []string `json:"dateList"`
	CountList []int    `json:"countList"`
}

type DataCountRequest struct {
	Type int `json:"type" form:"type"` // 0 七天内  1 一个月   2 一年
}

type DataLookDateRequest struct {
	DataCountRequest
	DocID uint `json:"docID" form:"docID"` // 文档id，不传就是查全部的
}

// DataLookDateView 文档浏览量数据
// @Tags 数据统计
// @Summary 文档浏览量数据
// @Description 文档浏览量数据
// @Param data query DataLookDateRequest true "参数"
// @Router /api/data/look_date [get]
// @Produce json
// @Success 200 {object} res.Response{data=DataCountResponse}
func (DataApi) DataLookDateView(c *gin.Context) {
	var cr DataLookDateRequest
	_ = c.ShouldBindQuery(&cr)

	var query = global.DB.Where("")
	now := time.Now()

	if cr.DocID != 0 {
		query.Where("docID = ?", cr.DocID)
	}

	var response DataCountResponse
	var dateTypeNum int

	switch cr.Type {
	case 1:
		// 一个月
		dateTypeNum = 30
	case 2:
		// 一年
		dateTypeNum = 365
	default:
		// 7天内
		dateTypeNum = 7
	}
	query.Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= createdAt", dateTypeNum))
	aDay := now.AddDate(0, 0, -dateTypeNum)
	for i := 1; i <= dateTypeNum; i++ {
		response.DateList = append(response.DateList, aDay.AddDate(0, 0, i).Format("2006-01-02"))
	}

	type dateCountType struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	var dateCountList []dateCountType
	global.DB.Model(models.DocDataModel{}).Where(query).
		Select(
			"date_format(createdAt, '%Y-%m-%d') as date",
			"sum(lookCount) as count").
		Group("date").Scan(&dateCountList)
	var dateCountMap = map[string]int{}
	for _, countType := range dateCountList {
		dateCountMap[countType.Date] = countType.Count
	}
	for _, s := range response.DateList {
		count, _ := dateCountMap[s]
		response.CountList = append(response.CountList, count)
	}
	res.OKWithData(response, c)
}
