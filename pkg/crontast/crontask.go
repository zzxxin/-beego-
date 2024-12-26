package crontast

import (
	"beegoweb/crons" // 导入 crons 包
	"beegoweb/models"
	"beegoweb/pkg/db"
	"log"
	"time"
)

// InitCronScheduler 初始化Cron调度器，定期检查任务状态并启动/停止任务
func InitCronScheduler() {
	// 同步方式启动任务
	crons.CronScheduler.Start() // 确保调度器已经启动

	// 每5秒检查一次总开关和任务状态
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// log.Println("循环检测")
		checkAndUpdateCronTasks()
	}
}

// checkAndUpdateCronTasks 检查总开关和任务状态，更新任务调度器
func checkAndUpdateCronTasks() {
	// 查询总开关状态
	var systemSetting models.SystemSetting
	if err := db.DB.Where("type = ?", "cron_status").First(&systemSetting).Error; err != nil {
		log.Println("无法查询总开关状态:", err)
		systemSetting.Status = 0 // 默认关闭
	}
	// 如果总开关关闭，停止所有任务
	if systemSetting.Status == 0 {
		crons.StopAllTasks()
		log.Println("总开关关闭，所有任务已停止")
		return
	}

	// 检查并处理每个任务的开关
	var tasks []models.CronTask
	if err := db.DB.Find(&tasks).Error; err != nil {
		log.Println("无法查询任务状态:", err)
		return
	}

	for _, task := range tasks {
		if task.TaskStatus == 1 { // 启用的任务
			crons.StartTask(task)
		} else { // 停用的任务
			crons.StopTask(task.TaskName)
		}
	}
}
