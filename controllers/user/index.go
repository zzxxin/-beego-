package user

import (
	"beegoweb/controllers"
	"beegoweb/models"
	services "beegoweb/service/user"
	"fmt"
	"strings"
	"time"
)

type UserController struct {
	controllers.BaseController
}

// 用户列表
func (u *UserController) UserList() {
	// 获取当前页码和每页显示数量
	page, _ := u.GetInt("page", 1)
	pageSize, _ := u.GetInt("pageSize", 10)

	// 调用服务层获取用户列表和分页信息
	users, pagination, err := services.GetUserList(page, pageSize)
	if err != nil {
		u.Alert("查询出错，请稍后重试", 500, "/error_page")
		return
	}

	// 将数据绑定到模板
	u.Data["user_list"] = users
	u.Data["Pagination"] = pagination
	u.TplName = "user/userlist.tpl"
}

// 添加用户
func (u *UserController) AddUser() {
	if u.Ctx.Input.IsPost() {
		userName := u.GetString("user_name")
		passwd := u.GetString("passwd")
		confirmpassword := u.GetString("confirm_password")
		realName := u.GetString("real_name")
		mobile := u.GetString("mobile")
		status, _ := u.GetInt("status")
		isSuper := u.GetString("is_super")

		if passwd != confirmpassword {
			u.Alert("密码不一致", 302, "/user_list")
			return
		}

		userService := &services.UserService{}
		msg, err := userService.AddUser(userName, passwd, realName, mobile, status, isSuper)
		if err != nil {
			u.Alert("添加失败【"+msg+"】请重试", 302, "/user_list")
			return
		}
		u.Alert("用户【"+userName+"】添加成功", 200, "/user_list")
	} else {
		u.Data["act"] = "add"
		u.TplName = "user/adduser.tpl"
	}
}

// 修改用户
func (u *UserController) UserEdit() {
	if u.Ctx.Request.Method == "POST" {
		userName := u.GetString("user_name")
		id, _ := u.GetInt("user_id")
		realName := u.GetString("real_name")
		mobile := u.GetString("mobile")
		status, _ := u.GetInt("status")
		isSuper := u.GetString("is_super")

		userService := &services.UserService{}
		msg, err := userService.UpdateUser(userName, id, realName, mobile, status, isSuper)
		if err != nil {
			u.Alert("操作失败"+msg+"请重试", 302, "/user_list")
			return
		}
		u.Alert("操作成功", 200, "/user_list")
	} else {
		id, _ := u.GetInt("id")
		userInfo, err := models.GetUserByID(uint(id))
		if err != nil || userInfo == nil {
			u.Redirect("/user_list?msg=用户不存在", 302)
			return
		}

		u.Data["user"] = userInfo
		u.Data["act"] = "update"
		u.Data["id"], _ = u.GetInt("id")
		u.TplName = "user/adduser.tpl"
	}
}

// 修改用户状态
func (u *UserController) UserStatus() {
	id, _ := u.GetInt("id", 0)
	status, _ := u.GetInt("status")

	if id == 0 {
		u.Alert("参数异常", 30001, "/user_list")
		return
	}
	fmt.Printf("修改用户ID为：%d, 修改状态为：%d\n", id, status)

	// 调用服务层修改用户状态
	err := services.UpdateUserStatus(int64(id), status)

	if err != nil {
		u.Alert("操作失败，请重试", 302, "/user_list")
	} else {
		u.Alert("操作成功", 200, "/user_list")
	}
}

// 用户数据导出
func (u *UserController) UserExport() {
	err := services.ExportUserData(u.Ctx.ResponseWriter)
	if err != nil {
		u.Ctx.WriteString("导出失败: " + err.Error())
	}
}

// 分配角色
func (u *UserController) AllotRole() {
	if u.Ctx.Input.IsPost() {
		bindChecked := u.GetStrings("bind_checked[]")
		userId, _ := u.GetInt("user_id")
		if userId == 0 {
			u.Alert("参数异常", 30001, "/user_list")
			return
		}

		// 处理角色ID
		var roleIds string
		if len(bindChecked) > 0 {
			roleIds = strings.Join(bindChecked, ",")
		}

		data := map[string]interface{}{
			"role_ids":    roleIds,
			"update_time": time.Now().Unix(),
		}

		fmt.Printf("修改用户内容为：%+v", data)

		// 更新用户角色
		if _, err := models.UpdateUserByQuery(data, "id = ?", userId); err != nil {
			u.Alert("操作失败", 30002, "/user_list")
			return
		}

		u.Alert("操作成功", 200, "/user_list")
		return
	}

	id, _ := u.GetInt("id")
	user, err := models.GetUserByID(uint(id))
	if err != nil || user == nil {
		u.Alert("用户信息异常", 3002, "/user_list")
		return
	}

	result := make(map[string]interface{})
	result["user_name"] = user.UserName
	result["user_id"] = user.ID
	result["bind_role"] = strings.Split(user.RoleIDs, ",")

	// 获取所有角色
	roleList, _ := models.GetRoleAllByQuery("1=1")
	fmt.Printf("角色列表: %+v\n", roleList)
	result["all_role"] = roleList

	u.Data["result"] = result
	u.TplName = "user/allotrole.tpl"

}
