// models/user.go
package models

import (
	"beegoweb/pkg/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type AdminUser struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserName   string    `gorm:"type:varchar(64);uniqueIndex:kfsid_username;not null" json:"user_name"`
	Passwd     string    `gorm:"type:varchar(64);not null" json:"passwd"`
	RealName   string    `gorm:"type:varchar(32);not null" json:"real_name"`
	Mobile     string    `gorm:"type:varchar(15);not null;index:mobile" json:"mobile"`
	Status     int       `gorm:"type:int(1);default:1;not null" json:"status"`
	IsSuper    string    `gorm:"type:char(1);default:'N';not null" json:"is_super"`
	Rights     string    `gorm:"type:text;not null" json:"rights"`
	RoleIDs    string    `gorm:"type:varchar(255);default:'';not null" json:"role_ids"`
	UpdateTime int       `gorm:"type:int;not null" json:"update_time"`
	CreateTime int       `gorm:"type:int;not null" json:"create_time"`
	LastLogin  time.Time `gorm:"type:datetime;not null" json:"last_login"`
	StatusName string    `gorm:"-" json:"status_name"` // 忽略此字段的数据库映射，依然可以在JSON中返回
}

// 设置表名
func (AdminUser) TableName() string {
	return "bk_admin_users"
}

func CreateUser(user *AdminUser) error {
	return db.DB.Create(user).Error
}

func GetUserByID(id uint) (*AdminUser, error) {
	var user AdminUser
	if err := db.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 修改用户信息
func UpdateUser(user *AdminUser, where string, args ...interface{}) error {
	// 只更新非零值字段
	return db.DB.Model(&AdminUser{}).Where(where, args...).Updates(user).Error
}

func UpdateFieldByID(id int64, fieldName string, fieldValue interface{}) error {
	return db.DB.Model(&AdminUser{}).Where("id = ?", id).UpdateColumn(fieldName, fieldValue).Error
}

// GetUserByQuery 根据条件查询用户
func GetUserByQuery(db *gorm.DB, where string, args ...interface{}) (*AdminUser, error) {
	var user AdminUser
	result := db.Where(where, args...).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

// GetUserByQueryPage 根据条件查询用户并分页
func GetUserByQueryPage(page int, pageSize int, where string, args ...interface{}) ([]AdminUser, int64, error) {
	var users []AdminUser
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询符合条件的总记录数
	err := db.DB.Model(&AdminUser{}).Where(where, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	result := db.DB.Where(where, args...).Order("id desc").Limit(pageSize).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return users, total, nil
}

// GetUserByQuery 根据条件查询用户
func GetUserQuery(where string, args ...interface{}) (*AdminUser, error) {
	var user AdminUser
	result := db.DB.Where(where, args...).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}

func GetUserAllByQuery(query string, args ...interface{}) ([]AdminUser, error) {
	var roles []AdminUser
	err := db.DB.Where(query, args...).Find(&roles).Error
	return roles, err
}

// 根据条件修改用户信息
func UpdateUserByQuery(fields map[string]interface{}, where string, args ...interface{}) (error, error) {
	// 使用 GORM 的 Updates 方法，只更新传递的字段
	return db.DB.Model(&AdminUser{}).Where(where, args...).Updates(fields).Error, nil
}
