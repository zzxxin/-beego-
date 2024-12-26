// models/user.go
package models

import (
	"beegoweb/pkg/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AdminRole struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	RoleName   string `gorm:"type:varchar(64);uniqueIndex:kfsid_username;not null" json:"role_name"`
	Rights     string `gorm:"type:text" json:"rights"`
	Operator   int64  `gorm:"type:int(11);not null" json:"operator"`
	UpdateTime int    `gorm:"type:int;not null" json:"update_time"`
	CreateTime int    `gorm:"type:int;not null" json:"create_time"`
}

// 设置表名
func (AdminRole) TableName() string {
	return "bk_admin_role"
}

func CreateRole(user *AdminRole) error {
	return db.DB.Create(user).Error
}

func GetRoleByID(id uint) (*AdminRole, error) {
	var role AdminRole
	if err := db.DB.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func UpdateRole(fields map[string]interface{}, where string, args ...interface{}) error {
	// 使用 GORM 的 Updates 方法，只更新传递的字段
	return db.DB.Model(&AdminRole{}).Where(where, args...).Updates(fields).Error
}

// GetUserByQuery 根据条件查询用户
func GetRoleByQuery(where string, args ...interface{}) (*AdminRole, error) {
	var role AdminRole
	result := db.DB.Where(where, args...).First(&role)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &role, result.Error
}

func GetRolesByIDs(ids []string, roles *[]AdminRole) error {
	result := db.DB.Where("id IN ?", ids).Find(roles)
	return result.Error
}

// GetRoleByQueryPage 根据条件查询用户并分页
func GetRoleByQueryPage(page int, pageSize int, where string, args ...interface{}) ([]AdminRole, int64, error) {
	var roles []AdminRole
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询符合条件的总记录数
	err := db.DB.Model(&AdminRole{}).Where(where, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	result := db.DB.Where(where, args...).Order("id desc").Limit(pageSize).Offset(offset).Find(&roles)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return roles, total, nil
}

func GetRoleAllByQuery(query string, args ...interface{}) ([]AdminRole, error) {
	var roles []AdminRole
	err := db.DB.Where(query, args...).Find(&roles).Error
	return roles, err
}
