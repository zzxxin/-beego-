package role

import (
	"beegoweb/models"
	"beegoweb/utils"
	"fmt"
	"time"
)

type RoleService struct{}

// GetUserList 获取角色列表并处理分页逻辑
func GetRoleList(page, pageSize int) ([]models.AdminRole, *utils.Pagination, error) {
	// 查询用户数据
	roles, total, err := models.GetRoleByQueryPage(page, pageSize, "1 = ? ", 1)
	if err != nil {
		return nil, nil, err
	}

	// 生成分页对象
	pagination := utils.NewPagination(int(total), pageSize, page)

	return roles, pagination, nil
}

// 添加角色
func AdRole(roleName string, operator int64) (string, error) {

	roleInfo, err := models.GetRoleByQuery("role_name = ?", roleName)
	if err != nil {
		return "数据查询失败", err
	}
	if roleInfo != nil {
		return "当前角色已添加", nil
	}

	roleAddData := &models.AdminRole{
		RoleName:   roleName,
		Operator:   operator,
		CreateTime: int(time.Now().Unix()),
	}

	fmt.Printf("添加数据为：%s", roleAddData)
	if err := models.CreateRole(roleAddData); err != nil {
		return "添加失败", err
	}
	return "角色添加成功", nil

}

// 修改角色
func Uprole(roleName string, operator int64, roleId int64) (string, error) {
	// 查询 角色信息
	roleInfo, err := models.GetRoleByID(uint(roleId))
	if err != nil {
		return "角色信息获取失败", err
	}

	if roleInfo.RoleName == roleName {
		return "角色名称不能为空", nil
	}
	userData := map[string]interface{}{
		"RoleName":   roleName,
		"UpdateTime": int(time.Now().Unix()),
		"Operator":   operator,
	}

	// 修改操作
	if err = models.UpdateRole(userData, "id = ?", roleId); err != nil {
		return "角色修改失败", err
	}
	return "角色修改成功", nil

}

// 获取权限列表分页获取
func GetRightList(page, pageSize int) ([]models.AdminRights, *utils.Pagination, error) {
	// 查询用户数据
	rightsList, total, err := models.GetRightByQueryPage(page, pageSize, "1 = ? ", 1)
	if err != nil {
		return nil, nil, err
	}

	// 生成分页对象
	pagination := utils.NewPagination(int(total), pageSize, page)

	return rightsList, pagination, nil
}
