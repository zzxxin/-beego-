package models

import (
	"time"
)

// Message 模型定义
type Message struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID string    `gorm:"index;not null" json:"channel_id"`
	UserID    uint      `gorm:"index;not null" json:"user_id"`
	Username  string    `gorm:"not null" json:"username"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"` // 自动生成时间戳
}

// 定义表名
func (Message) TableName() string {
	return "bk_scoket_messages"
}
