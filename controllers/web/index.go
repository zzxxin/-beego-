package web

import (
	"beegoweb/controllers"
	"beegoweb/models"
	"beegoweb/utils"
)

type IndexController struct {
	controllers.BaseController
}

// Index 首页
// Index 首页
func (c *IndexController) Index() {
	if !c.CheckLogin() {
		return // 用户未登录，返回已处理的 JSON 响应
	}

	// 获取用户信息
	userId := c.Ctx.Input.GetData("userId").(uint)
	userInfo, err := models.GetUserByID(userId)
	if err != nil {
		utils.JSONResponse(&c.Controller, 401, "账号异常", nil)
		return
	}

	// 渲染模板
	c.RenderTemplate("web/index.tpl", map[string]interface{}{
		"user_info": userInfo,
	})
}

func (c *IndexController) Webmain() {
	c.TplName = "web/main.tpl"
}
