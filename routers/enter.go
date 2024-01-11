package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvd_server/global"
	"gvd_server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	router := gin.Default()

	if global.Config.System.Env == "dev" {
		router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}

	// 创建了一个以api开头的api分组
	apiGroup := router.Group("api")
	apiGroup.Use(middleware.LogMiddleWare())
	// 又将这个api分组赋给了 RouterGroup
	routerGroup := RouterGroup{apiGroup}

	// 线上如果有nginx，这一步可以省略  第一个参数是web的访问别名  第二个参数是内部的映射目录
	router.Static("/uploads", "uploads")

	routerGroup.UserRouter()
	//routerGroup.ImageRouter()
	//routerGroup.LogRouter()
	//routerGroup.SiteRouter()
	//routerGroup.RoleRouter()
	//routerGroup.DocRouter()
	//routerGroup.RoleDocRouter()
	//routerGroup.DataRouter()
	//routerGroup.UserCenterRouter()
	return router
}
