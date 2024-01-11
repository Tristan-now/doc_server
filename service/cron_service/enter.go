package cron_service

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CornInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")

	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))

	// 每天的2点去同步数据
	Cron.AddFunc("0 0 2 * * ?", SyncDocData)
	// 每天的3点去同步文档的浏览数据
	Cron.AddFunc("0 0 3 * * ?", SyncDocDataDate)

	Cron.Start()
}
