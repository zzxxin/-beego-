package models

import (
	"beegoweb/pkg/db"
	"errors"

	"gorm.io/gorm"
)

type SystemSetting struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Type         string `gorm:"type:varchar(64);not null" json:"type"`
	Status       int    `gorm:"type:int(1);not null" json:"status"`
	ErrorMessage string `gorm:"type:varchar(255)"` // 记录错误信息
	UpdateTime   int    `gorm:"type:int;not null" json:"update_time"`
	CreateTime   int    `gorm:"type:int;not null" json:"create_time"`
}

// 定义表名
func (SystemSetting) TableName() string {
	return "bk_system_setting"
}

// 创建配置
func CreateSetting(sysSet *SystemSetting) error {
	return db.DB.Create(sysSet).Error
}

// 修改配置
func UpdateSetting(fields map[string]interface{}, where string, args ...interface{}) error {
	// 使用 GORM 的 Updates 方法，只更新传递的字段
	return db.DB.Model(&SystemSetting{}).Where(where, args...).Updates(fields).Error
}

// 获取配置
func GetSettingByQuery(where string, args ...interface{}) (*SystemSetting, error) {
	var set SystemSetting
	result := db.DB.Where(where, args...).First(&set)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &set, result.Error
}
