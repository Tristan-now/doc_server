package site_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/config"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/service/common/res"
	"reflect"
)

func (SiteApi) SiteUpdateView(c *gin.Context) {
	var cr config.Site
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	updateStructValue(cr, reflect.ValueOf(&global.Config.Site))

	core.SetYaml()

	res.OKWithMsg("更新成功", c)
}

func updateStructValue(data any, oldValue reflect.Value) {
	v := reflect.ValueOf(data)
	var updateIndexSlice []int
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			// 如果不为空，就加入到要更新的列表里
			updateIndexSlice = append(updateIndexSlice, i)
		}
	}

	t := reflect.TypeOf(data)
	for _, updateIndex := range updateIndexSlice {
		// 拿字段名
		name := t.Field(updateIndex).Name
		// 拿到字段的value
		fild := v.Field(updateIndex)
		// 动态修改
		oldValue.Elem().FieldByName(name).Set(fild)
	}
}
