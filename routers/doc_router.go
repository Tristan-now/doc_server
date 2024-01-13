package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) DocRouter() {
	app := api.App.DocApi
	r := router.Group("docs")
	r.POST("", middleware.JwtAdmin(), app.DocCreateView)
	r.GET("info/:id", middleware.JwtAdmin(), app.DocInfoView)
	r.GET(":id", app.DocContentView)
}
