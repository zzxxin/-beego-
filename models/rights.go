package models

import (
	"beegoweb/pkg/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type AdminRights struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;column:id"`            // 主键ID
	GroupName  string `gorm:"type:varchar(30);default:'';column:group_name"` // 分组名称
	RightName  string `gorm:"type:varchar(30);default:'';column:right_name"` // 权限名称
	RightLogo  string `gorm:"type:varchar(30);default:'';column:right_logo"` // 权限标识
	CreateTime int    `gorm:"type:int;default:0;column:create_time"`         // 创建时间
	UpdateTime int    `gorm:"type:int;default:0;column:update_time"`         // 修改时间
	OperatorID int    `gorm:"type:int;default:0;column:operator_id"`         // 操作人
	IsRight    int    `gorm:"type:tinyint(1);default:1;column:is_right"`     // 是否校验权限
	IsMenu     int    `gorm:"type:tinyint(1);default:1;column:is_menu"`      // 是否菜单
}

// 设置表名
func (AdminRights) TableName() string {
	return "bk_rights"
}

func CreateRight(user *AdminRights) error {
	return db.DB.Create(user).Error
}

func GetRightByID(id uint) (*AdminRights, error) {
	var right AdminRights
	if err := db.DB.First(&right, id).Error; err != nil {
		return nil, err
	}
	return &right, nil
}

func UpdateRight(fields map[string]interface{}, where string, args ...interface{}) error {
	// 使用 GORM 的 Updates 方法，只更新传递的字段
	return db.DB.Model(&AdminRights{}).Where(where, args...).Updates(fields).Error
}

// GetUserByQuery 根据条件查询用户
func GetRightByQuery(where string, args ...interface{}) (*AdminRights, error) {
	var right AdminRights
	result := db.DB.Where(where, args...).First(&right)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &right, result.Error
}

func GetAllRights(where string, args ...interface{}) ([]AdminRights, error) {
	var rights []AdminRights
	result := db.DB.Where(where, args...).Find(&rights)
	return rights, result.Error
}

func GetRightsByIDs(ids []uint) ([]AdminRights, error) {
	var rights []AdminRights
	result := db.DB.Where("id IN ? AND is_menu = ?", ids, 1).Find(&rights)
	return rights, result.Error
}

// 获取角色列表
func GetRightPriList(query string, args ...interface{}) ([]AdminRights, error) {
	var adminrights []AdminRights
	err := db.DB.Where(query, args...).Find(&adminrights).Error
	return adminrights, err
}

func GetRightByQueryPage(page int, pageSize int, where string, args ...interface{}) ([]AdminRights, int64, error) {
	var rights []AdminRights
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询符合条件的总记录数
	err := db.DB.Model(&AdminRights{}).Where(where, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	result := db.DB.Where(where, args...).Order("id desc").Limit(pageSize).Offset(offset).Find(&rights)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return rights, total, nil
}
