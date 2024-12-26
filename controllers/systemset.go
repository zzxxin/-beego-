package controllers

import (
	services "beegoweb/service"
)

type SystemSetController struct {
	BaseController
}

/*
 * 设置定时任务开关
 */
func (u *SystemSetController) CronSet() {
	Status, _ := u.GetInt("status")
	Type := u.GetString("type")
	if Type == "" {
		u.Alert("请检查提交参数", 301, "/cron_list")
		return
	}

	// 查询是否已经添加过
	systemService := &services.SystemSetService{}
	_, err := systemService.HandleSystemSet(Type, Status)

	if err != nil {
		u.Alert("操作失败，请重试", 302, "/cron_list")
	} else {
		u.Alert("操作成功", 200, "/cron_list")
	}
}
