package role

import (
	"beegoweb/controllers"
	"beegoweb/models"
	services "beegoweb/service"
	"beegoweb/service/role"
	"fmt"
	"time"
)

type RightController struct {
	controllers.BaseController
}

// 权限列表
func (c *RightController) RightList() {
	// 获取当前页码和每页显示数量
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 10)

	// 调用服务层获取用户列表和分页信息
	rightList, pagination, err := role.GetRightList(page, pageSize)
	if err != nil {
		c.Alert("查询出错，请稍后重试", 500, "/error_page")
		return
	}

	// 将数据绑定到模板
	c.Data["right_list"] = rightList
	c.Data["Pagination"] = pagination
	c.TplName = "role/rightlist.tpl"
}

func (c *RightController) AddRight() {

	if c.Ctx.Input.IsPost() {
		groupName := c.GetString("group_name")
		rightLogo := c.GetString("right_logo")
		rightName := c.GetString("right_name")
		isRight, _ := c.GetInt("is_right")
		isMenu, _ := c.GetInt("is_menu")

		// 参数校验
		if groupName == "" || rightLogo == "" || rightName == "" {
			c.Alert("参数异常", 30002, "/right_list")
			return
		}

		// 校验当前路由是否已经添加
		rightInfo, _ := models.GetRightByQuery("right_logo = ?", rightLogo)
		if rightInfo != nil {
			c.Alert("当前权限已添加,请重新配置", 3001, "/right_list")
			return
		}

		var operator uint
		if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
			operator = op
		}
		// 数据准备
		right := &models.AdminRights{
			GroupName:  groupName,
			RightLogo:  rightLogo,
			RightName:  rightName,
			IsRight:    isRight,
			IsMenu:     isMenu,
			CreateTime: int(time.Now().Unix()),
			OperatorID: int(operator), // 根据实际情况修改
		}

		// 插入数据
		if err := models.CreateRight(right); err != nil {
			c.Alert("添加权限失败", 3001, "/right_list")
			return
		}

		// 成功返回
		c.Alert("添加权限成功", 200, "/right_list")
		return
	}

	c.Data["act"] = "add"
	c.TplName = "role/addright.tpl"

}

func (c *RightController) UpRight() {

	if c.Ctx.Input.IsPost() {
		id, _ := c.GetInt("right_id")
		groupName := c.GetString("group_name")
		rightLogo := c.GetString("right_logo")
		rightName := c.GetString("right_name")
		isRight, _ := c.GetInt("is_right")
		isMenu, _ := c.GetInt("is_menu")

		// 参数校验
		if groupName == "" || rightLogo == "" || rightName == "" || isRight == 0 || isMenu == 0 {
			c.Alert("参数异常", 30002, "/right_list")
			return
		}

		var operator uint
		if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
			operator = op
		}

		// 数据准备
		upRight := map[string]interface{}{
			"GroupName":  groupName,
			"RightLogo":  rightLogo,
			"RightName":  rightName,
			"IsRight":    isRight,
			"IsMenu":     isMenu,
			"UpdateTime": int(time.Now().Unix()),
			"OperatorID": int(operator), // 根据实际情况修改
		}
		// 插入数据
		if err := models.UpdateRight(upRight, "id = ?", id); err != nil {
			c.Alert("修改权限失败", 3001, "/right_list")
			return
		}
		log := services.AddOperateLog(c.Ctx, "编辑权限", "cfasd", 1)
		fmt.Printf("日志记录情况%s", log)

		// 成功返回
		c.Alert("修改权限成功", 200, "/right_list")
		return

	}

	id, _ := c.GetInt("id")
	// 获取权限信息
	rightInfo, _ := models.GetRightByID(uint(id))
	c.Data["right_info"] = rightInfo
	c.Data["act"] = "update"
	c.TplName = "role/addright.tpl"
}
