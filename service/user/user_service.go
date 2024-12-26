package services

import (
	"beegoweb/models"
	"beegoweb/utils"
	"net/http"
	"time"
)

type UserService struct{}

// GetUserList 获取用户列表并处理分页逻辑
func GetUserList(page, pageSize int) ([]models.AdminUser, *utils.Pagination, error) {
	// 查询用户数据
	users, total, err := models.GetUserByQueryPage(page, pageSize, "1 = ? ", 1)
	if err != nil {
		return nil, nil, err
	}

	// 转换状态为中文
	for i := range users {
		switch users[i].Status {
		case 1:
			users[i].StatusName = "启用"
		default:
			users[i].StatusName = "停用"
		}

		switch users[i].IsSuper {
		case "Y":
			users[i].IsSuper = "是"
		case "N":
			users[i].IsSuper = "否"
		}
	}

	// 生成分页对象
	pagination := utils.NewPagination(int(total), pageSize, page)

	return users, pagination, nil
}

// UpdateUserStatus 修改用户状态
func UpdateUserStatus(id int64, status int) error {
	// 调用模型层更新字段
	return models.UpdateFieldByID(id, "status", status)
}

// ExportUserData 查询用户并导出为 Excel 文件
func ExportUserData(writer http.ResponseWriter) error {
	// 表头
	headers := []string{"用户ID", "姓名", "密码", "状态", "是否超管"}

	// 查询用户数据
	users, _, err := models.GetUserByQueryPage(1, 10, "1 = ? ", 1)
	if err != nil {
		return err
	}

	// 转换用户数据
	var data [][]interface{}
	for _, user := range users {
		// 转换状态为中文
		var statusName string
		switch user.Status {
		case 1:
			statusName = "启用"
		default:
			statusName = "停用"
		}

		// 转换是否为超管
		var isSuper string
		switch user.IsSuper {
		case "Y":
			isSuper = "是"
		case "N":
			isSuper = "否"
		}

		// 填充数据
		data = append(data, []interface{}{
			user.ID,
			user.UserName,
			user.Passwd,
			statusName,
			isSuper,
		})
	}

	// 调用工具类生成 Excel
	f, err := utils.ExportExcel("Sheet1", headers, data)
	if err != nil {
		return err
	}

	// 设置文件名
	filename := "用户列表" + time.Now().Format("20060102_150405") + ".xlsx"

	// 将文件写入 HTTP 响应
	return utils.SaveExcelToResponse(f, filename, writer)
}

func (s *UserService) UpdateUser(userName string, id int, realName string, mobile string, status int, isSuper string) (string, error) {
	// 验证必填字段
	if userName == "" || realName == "" || mobile == "" {
		return "请完善信息", nil
	}

	// 检查用户名是否已存在
	userInfo, _ := models.GetUserByID(uint(id))
	if userInfo != nil && userInfo.ID != uint(id) {
		return "登录用户名已被其他用户使用，请修改", nil
	}

	// 更新用户信息
	user := models.AdminUser{
		ID:         uint(id),
		UserName:   userName,
		RealName:   realName,
		Mobile:     mobile,
		Status:     status,
		IsSuper:    isSuper,
		UpdateTime: int(time.Now().Unix()),
	}
	if err := models.UpdateUser(&user, "id = ?", uint(id)); err != nil {
		return "用户修改失败", err
	}

	return "修改成功", nil
}

// 添加用户数据
func (s *UserService) AddUser(userName string, passwd string, realName string, mobile string, status int, isSuper string) (string, error) {

	// 验证必填字段
	if userName == "" || realName == "" || mobile == "" || passwd == "" {
		return "请完善信息", nil
	}

	// 检查用户名是否已存在
	userInfo, _ := models.GetUserQuery("user_name = ?", userName)
	if userInfo != nil {
		return "登录用户名已被其他用户使用，请修改", nil
	}

	// 新增用户信息
	user := models.AdminUser{
		Passwd:     utils.GenerateMD5Hash(passwd),
		UserName:   userName,
		RealName:   realName,
		Mobile:     mobile,
		Status:     status,
		IsSuper:    isSuper,
		CreateTime: int(time.Now().Unix()),
	}
	if err := models.CreateUser(&user); err != nil {
		return "用户添加失败", err
	}
	return "用户添加成功", nil
}
