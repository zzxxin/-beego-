package models

// Channel 模型定义
type Channel struct {
	ID      string `gorm:"primaryKey" json:"id"`                  // 主键，频道 ID
	Name    string `gorm:"not null" json:"name"`                  // 频道名称
	IsGroup bool   `gorm:"not null;default:true" json:"is_group"` // 是否是群聊
	Members string `gorm:"not null" json:"members"`               // 成员列表，存储为 JSON
}

// 定义表名
func (Channel) TableName() string {
	return "bk_socket_channels"
}
