package main

import (
	"beegoweb/crons"
	"beegoweb/crons/jobs"
	"beegoweb/pkg"
	"beegoweb/pkg/crontast"
	"beegoweb/pkg/db"
	"beegoweb/pkg/redis"
	_ "beegoweb/routers"
	services "beegoweb/service"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.SetStaticPath("/static", "static")
	// 启用session会话
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.BConfig.RunMode = "dev"
	beego.BConfig.Log.AccessLogs = true

	// 初始化数据库连接
	db.Init()
	// 初始化redis and Mq链接
	redis.Init()
	// 工具以及定时任务初始化
	pkg.Init()

	// 注册定时任务 不用每次都去添加switch
	crons.RegisterTask("print_number", jobs.PrintNumberTask)
	crons.RegisterTask("print_data", jobs.PrintNumberDataTask)
	go func() {
		// 启动调度器并初始化任务 放到子进程中
		crontast.InitCronScheduler()
	}()

	// 链接 MQ
	services.InitRabbitMQ()
	defer services.CloseRabbitMQ()
	beego.Run()

}
