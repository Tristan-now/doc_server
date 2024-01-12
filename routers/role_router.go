package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleRouter() {
	app := api.App.RoleApi
	r := router.Group("roles").Use(middleware.JwtAdmin())
	r.POST("", app.RoleCreateView)   //角色创建
	r.GET("", app.RoleListView)      //角色列表
	r.PUT("", app.RoleUpdateView)    //角色更新
	r.DELETE("", app.RoleRemoveView) //角色更新
}
