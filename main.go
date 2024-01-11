package main

import (
	"gvd_server/core"
	_ "gvd_server/docs"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
	"gvd_server/service/cron_service"
)

// @title 文档项目api文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8082
// @BasePath /
func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)
	global.ESClient = core.InitEs()
	global.AddrDB = core.InitAddrDB()

	option := flags.Parse()
	if option.Run() {
		return
	}

	cron_service.CornInit()

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
