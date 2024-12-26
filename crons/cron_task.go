package crons

import (
	"beegoweb/models"
	"log"
	"sync"

	"github.com/robfig/cron/v3"
)

// CronScheduler 定时任务调度器
var CronScheduler *cron.Cron
var taskEntries map[string]cron.EntryID
var mu sync.Mutex

// 初始化调度器
func init() {
	CronScheduler = cron.New(cron.WithSeconds()) // 启用秒级调度
	taskEntries = make(map[string]cron.EntryID)
	CronScheduler.Start()
	log.Println("Cron调度器已初始化并启动")
}

// 启动任务
func StartTask(task models.CronTask) {
	mu.Lock()
	defer mu.Unlock()

	// 检查任务是否已经存在
	if _, exists := taskEntries[task.TaskName]; exists {
		// log.Printf("任务 [%s] 已经存在，跳过添加\n", task.TaskName)
		return
	}

	entryID, err := CronScheduler.AddFunc(task.CronExpression, func() {
		go ExecuteTask(task.TaskName) // 动态执行任务
	})
	if err != nil {
		// log.Printf("任务 [%s] 添加失败：%s\n", task.TaskName, err)
		return
	}

	taskEntries[task.TaskName] = entryID
	// log.Printf("任务 [%s] 已成功添加到调度器，EntryID: %d\n", task.TaskName, entryID)
}

// 停止任务
func StopTask(taskName string) {
	mu.Lock()
	defer mu.Unlock()

	// 如果任务不存在，直接返回
	entryID, exists := taskEntries[taskName]
	if !exists {
		return
	}

	// 移除任务
	CronScheduler.Remove(entryID)
	delete(taskEntries, taskName)
	log.Printf("任务 [%s] 已停止\n", taskName)
}

// 停止所有任务
func StopAllTasks() {
	mu.Lock()
	defer mu.Unlock()

	// 停止所有任务
	for taskName := range taskEntries {
		CronScheduler.Remove(taskEntries[taskName])
		delete(taskEntries, taskName)
	}
	log.Println("所有任务已停止")
}

// 定义任务处理函数类型
type TaskFunc func()

// 注册表
var taskRegistry = make(map[string]TaskFunc)

// 注册任务到任务注册表
func RegisterTask(taskName string, taskFunc TaskFunc) {
	if _, exists := taskRegistry[taskName]; exists {
		log.Printf("任务 [%s] 已经注册，跳过注册\n", taskName)
		return
	}
	taskRegistry[taskName] = taskFunc
	log.Printf("任务 [%s] 注册成功\n", taskName)
}

// 执行任务
func ExecuteTask(taskName string) {
	if taskFunc, exists := taskRegistry[taskName]; exists {
		taskFunc() // 执行任务函数
	} else {
		log.Printf("任务 [%s] 未注册，无法执行\n", taskName)
	}
}

// // 执行具体任务
// func ExecuteTask(taskName string) {
// 	log.Printf("执行任务: %s\n", taskName)
//
// 	// 根据任务名称调用实际的任务方法
// 	switch taskName {
// 	case "print_number":
// 		// 在这里调用实际的任务函数
// 		jobs.PrintNumberTask()
// 	case "print_data":
// 		// 在这里调用实际的任务函数
// 		jobs.PrintNumberDataTask()
// 	default:
// 		log.Println("未知任务:", taskName)
// 	}
// }
