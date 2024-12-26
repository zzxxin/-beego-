package models

import "beegoweb/pkg/db"

type OperatorLog struct {
	ID         uint   `gorm:"primaryKey"`                   // 主键
	Desc       string `gorm:"type:varchar(20);not null"`    // 操作类型
	Operator   string `gorm:"type:varchar(50);not null"`    // 操作人ID
	Request    string `gorm:"type:text"`                    // 请求参数
	Response   string `gorm:"type:text"`                    // 返回结果
	Router     string `gorm:"type:varchar(255);not null"`   // 路由地址
	Method     string `gorm:"type:varchar(255);not null"`   // 方法名
	CreateTime int64  `gorm:"autoCreateTime:nano;not null"` // 创建时间（Unix 时间戳）
	Type       int    `gorm:"type:tinyint(1);default:1"`    // 日志类型 1后台 2前台
}

// 定义表名
func (OperatorLog) TableName() string {
	return "bk_operate_log"
}

func CreateOperatorLog(user *OperatorLog) error {
	return db.DB.Create(user).Error
}
