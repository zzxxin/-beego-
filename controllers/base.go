package controllers

import (
	"beegoweb/pkg/db"
	"beegoweb/utils"

	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

type BaseController struct {
	beego.Controller
	// 数据库实例
	DB *gorm.DB
}

func (c *BaseController) Prepare() {
	// 从配置文件中读取 BaseUrl
	baseUrl, _ := beego.AppConfig.String("js_base_url")
	// 设置全局模板变量
	c.Data["BaseUrl"] = baseUrl
	// 使用全局数据库连接
	c.DB = db.DB

	// 添加用户权限数据
	c.Data["UserRight"] = c.Ctx.Input.GetData("userRight")
}

func (c *BaseController) Alert(message string, code int, url string) {
	c.Data["message"] = message
	c.Data["code"] = code
	c.Data["url"] = url
	c.TplName = "common/alert.tpl"
}

// RenderTemplate 用于渲染模板并统一处理模板数据
func (c *BaseController) RenderTemplate(templateName string, data map[string]interface{}) {
	// 设置模板路径
	c.TplName = templateName

	// 设置模板数据
	for key, value := range data {
		c.Data[key] = value
	}
}

// 自定义方法：检查用户登录状态并返回错误信息
func (c *BaseController) CheckLogin() bool {
	userId := c.Ctx.Input.GetData("userId")
	if userId == nil {
		utils.JSONResponse(&c.Controller, 401, "用户未登录", nil)
		return false
	}
	return true
}
