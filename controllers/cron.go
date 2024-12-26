package controllers

import (
	"beegoweb/models"
	"beegoweb/utils"
	"fmt"
	"time"
)

type CronController struct {
	BaseController
}

// 计划任务列表
func (c *CronController) CronList() {

	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 10)

	// 查询用户数据
	cronList, total, _ := models.GetCronByQueryPage(page, pageSize, "1 = ? ", 1)

	// 生成分页对象
	pagination := utils.NewPagination(int(total), pageSize, page)

	// 获取cron 开关
	cronSetInfo, _ := models.GetSettingByQuery("type = ?", "cron_status")
	cronStatus := 0
	if cronSetInfo == nil {
		cronStatus = 0
	} else {
		cronStatus = cronSetInfo.Status
	}

	c.Data["cron_status"] = cronStatus
	// 将数据绑定到模板
	c.Data["cron_list"] = cronList
	c.Data["Pagination"] = pagination
	c.TplName = "cron/cronlist.tpl"

}

// 添加计划任务
func (c *CronController) AddCron() {
	if c.Ctx.Input.IsPost() {
		taskName := c.GetString("task_name")
		cronExpression := c.GetString("cron_expression")
		taskStatus, _ := c.GetInt("task_status")
		taskDesc := c.GetString("task_desc")
		if taskName == "" || cronExpression == "" || taskStatus == 0 || taskDesc == "" {
			c.Alert("请检查提交参数", 301, "/cron_list")
			return
		}

		// 查询是否已经添加过
		cronInfo, _ := models.GetCronByQuery("task_name = ?", taskName)
		if cronInfo != nil {
			c.Alert("当前任务已经添加", 302, "/cron_list")
			return
		}

		addData := &models.CronTask{
			TaskName:       taskName,
			CronExpression: cronExpression,
			TaskStatus:     int(taskStatus),
			TaskDesc:       taskDesc,
			CreatedAt:      time.Time{},
		}
		if err := models.CreateCron(addData); err != nil {
			c.Alert("任务添加失败", 301, "/cron_list")
			return
		}
		c.Alert("任务添加成功", 200, "/cron_list")
		return
	}

	c.Data["act"] = "add"
	c.TplName = "cron/addcron.tpl"
}

// 修改cron
func (c *CronController) UpCron() {
	if c.Ctx.Input.IsPost() {
		cronId, _ := c.GetInt("cron_id")
		taskName := c.GetString("task_name")
		cronExpression := c.GetString("cron_expression")
		taskStatus, _ := c.GetInt("task_status")
		taskDesc := c.GetString("task_desc")
		if taskName == "" || cronExpression == "" || taskStatus == 0 || taskDesc == "" {
			c.Alert("请检查提交参数", 301, "/cron_list")
			return
		}

		upData := map[string]interface{}{
			"TaskName":       taskName,
			"CronExpression": cronExpression,
			"TaskStatus":     int(taskStatus),
			"TaskDesc":       taskDesc,
			"UpdatedAt":      time.Time{},
		}
		fmt.Printf("修改数据为：%+v \n", upData)
		fmt.Printf("修改ID：%d", cronId)
		if err := models.UpdateCron(upData, "id = ?", cronId); err != nil {
			c.Alert("任务修改失败", 301, "/cron_list")
			return
		}
		c.Alert("任务修改成功", 200, "/cron_list")
		return
	}

	id, _ := c.GetInt("id")
	fmt.Printf("ID为：%d", id)
	cronInfo, _ := models.GetCronByQuery("id =?", id)

	c.Data["cron_info"] = cronInfo
	c.Data["act"] = "update"
	c.TplName = "cron/addcron.tpl"
}

// 修改cron状态
func (c *CronController) UpCronStatus() {
	id, _ := c.GetInt("id", 0)
	status, _ := c.GetInt("status")

	if id == 0 {
		c.Alert("参数异常", 30001, "/user_list")
		return
	}

	// 调用服务层修改用户状态
	err := models.UpdateCronStatus(int64(id), "task_status", status)

	if err != nil {
		c.Alert("操作失败，请重试", 302, "/cron_list")
	} else {
		c.Alert("操作成功", 200, "/cron_list")
	}
}
