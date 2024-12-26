package role

import (
	"beegoweb/controllers"
	"beegoweb/models"
	"beegoweb/service/role"
	"encoding/json"
	"fmt"
	"time"
)

type RoleController struct {
	controllers.BaseController
}

// 权限列表
func (c *RoleController) RoleList() {

	// 测试查询
	query := "id = ?"
	args := []interface{}{1}
	users, err := models.GetUserAllByQuery(query, args...)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 打印结果
	fmt.Printf("Users: %v\n", users)

	// 获取当前页码和每页显示数量
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 5)

	roles, pagination, err := role.GetRoleList(page, pageSize)
	if err != nil {
		c.Alert("查询出错，请稍后重试", 500, "/error_page")
		return
	}

	// 将数据绑定到模板
	c.Data["role_list"] = roles
	c.Data["Pagination"] = pagination
	c.TplName = "role/rolelist.tpl"
}

// 添加角色
func (c *RoleController) AddRole() {
	if c.Ctx.Request.Method == "POST" {
		roleNmae := c.GetString("role_name")
		if roleNmae == "" {
			c.Alert("请填写角色名称", 302, "")
			return
		}
		var operator uint
		if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
			operator = op
		}

		msg, err := role.AdRole(roleNmae, int64(operator))
		if err != nil {
			c.Alert("添加失败【"+msg+"】请重试", 302, "/role_list")
			return
		}
		c.Alert("角色【"+roleNmae+"】添加成功", 200, "/role_list")
	} else {
		c.Data["act"] = "add"
		c.TplName = "role/addrole.tpl"
	}
}

// 修改角色内容
func (c *RoleController) UpRole() {
	if c.Ctx.Request.Method == "POST" {
		roleNmae := c.GetString("role_name")
		roleId, _ := c.GetInt("role_id")
		if roleNmae == "" {
			c.Alert("请填写角色名称", 302, "")
			return
		}

		var operator uint
		if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
			operator = op
		}

		msg, err := role.Uprole(roleNmae, int64(operator), int64(roleId))
		if err != nil {
			c.Alert("角色修改失败【"+msg+"】请重试", 302, "/role_list")
			return
		}
		c.Alert("角色【"+roleNmae+"】修改成功", 200, "/role_list")
	} else {
		// 查询角色信息
		roleId, _ := c.GetInt("id")
		roleInfo, _ := models.GetRoleByID(uint(roleId))
		c.Data["roleinfo"] = roleInfo
		c.Data["act"] = "update"
		c.TplName = "role/addrole.tpl"
	}
}

// 角色分配权限
func (c *RoleController) RoleBindRight() {
	if c.Ctx.Input.IsPost() {
		roleId := c.GetString("role_id")
		bindChecked := c.GetStrings("bind_checked[]")
		fmt.Printf("拿到的绑定ID为：%+v", bindChecked)
		bindData, err := json.Marshal(bindChecked)
		if err != nil {
			c.Alert("操作失败", 3002, "/role_list")
			return
		}
		// 构建更新数据
		upData := map[string]interface{}{
			"Rights":     bindData,
			"UpdateTime": time.Now().Unix(), // 更新当前时间
		}
		fmt.Printf("保存内容为%+v", upData)
		err = models.UpdateRole(upData, "id = ? ", roleId)
		if err != nil {
			c.Alert("操作失败", 30004, "/role_list")
			return
		}
		c.Alert("操作成功", 200, "/role_list")
		return
	}

	id, err := c.GetInt("id")
	if err != nil {
		c.Alert("请选择用户", 3001, "/role_list")
		return
	}
	// 获取角色信息
	roleInfo, err := models.GetRoleByID(uint(id))
	if err != nil {
		c.Alert("角色信息异常", 3002, "/role_list")
		return
	}

	rightList, err := models.GetRightPriList("is_right = ?", 1)
	if err != nil {
		c.Alert("权限信息获取失败", 3003, "/role_list")
		return
	}

	var roleBind []string
	if roleInfo.Rights != "" {
		err = json.Unmarshal([]byte(roleInfo.Rights), &roleBind)
		if err != nil {
			c.Alert("角色数据处理失败", 3003, "/role_list")
			return
		}
	}

	result := map[string]interface{}{
		"role_name":  roleInfo.RoleName,
		"role_id":    roleInfo.ID,
		"bind_right": roleBind,
	}

	right := make(map[string][]models.AdminRights)
	for _, val := range rightList {
		groupName := val.GroupName
		if _, exists := right[groupName]; !exists {
			right[groupName] = []models.AdminRights{}
		}
		right[groupName] = append(right[groupName], val)
	}

	result["right_list"] = right
	c.Data["result"] = result
	c.TplName = "role/bindright.tpl"

}
