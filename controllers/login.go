package controllers

import (
	"beegoweb/models"
	"beegoweb/utils"
	"fmt"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	BaseController
}

// UserLogin 用户登录页
func (c *LoginController) UserLogin() {
	c.TplName = "user/login.tpl"
}

type User struct {
	ID   uint
	Name string
}

// 用户登录操作
func (c *LoginController) DoLogin() {
	UserName := c.GetString("user_name")
	Passwd := c.GetString("passwd")
	// 查询用户
	userInfo, err := models.GetUserByQuery(c.DB, fmt.Sprintf("user_name = '%s'", UserName))
	if err != nil || userInfo == nil {
		c.Data["json"] = map[string]interface{}{
			"code": 401,
			"msg":  "用户不存在",
		}
		if err := c.ServeJSON(); err != nil {
			c.Ctx.WriteString("Error serving JSON: " + err.Error())
			return
		}
		return
	}

	// 验证密码
	if !validatePassword(userInfo.Passwd, Passwd) {
		c.Data["json"] = map[string]interface{}{
			"code": 401,
			"msg":  "密码错误",
		}
		if err := c.ServeJSON(); err != nil {
			c.Ctx.WriteString("Error serving JSON: " + err.Error())
			return
		}
		return
	}

	// 生成 JWT
	jwtToken, err := utils.GenerateToken(userInfo.ID, userInfo.UserName)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"code": 500,
			"msg":  "生成 token 失败",
		}
		if err := c.ServeJSON(); err != nil {
			c.Ctx.WriteString("Error serving JSON: " + err.Error())
			return
		}
		return
	}
	user_login, _ := web.AppConfig.String("login_identification")
	// 设置 Cookie
	c.Ctx.SetCookie(user_login, jwtToken, 24*time.Hour, "/", "", false, true)

	// 返回成功信息
	utils.JSONResponse(&c.Controller, 200, "操作成功", map[string]interface{}{
		"token":     jwtToken,
		"uid":       userInfo.ID,
		"user_name": userInfo.UserName,
	})
}

// validatePassword 验证密码
func validatePassword(storedPassword, providedPassword string) bool {
	// 在这里使用你的密码验证逻辑，例如使用 bcrypt 进行密码验证
	return storedPassword == utils.GenerateMD5Hash(providedPassword)
}

// 退出登录
func (c *LoginController) LoginOut() {
	// 假设你用的Cookie名称为"login_identification"
	user_login, _ := web.AppConfig.String("login_identification")
	// 清除 Cookie
	c.Ctx.SetCookie(user_login, "", -1, "/")

	// 重定向到登录页面或其他页面
	c.Redirect("/user_login", 302)
}
