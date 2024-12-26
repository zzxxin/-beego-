package services

import (
	"beegoweb/models"
	"beegoweb/pkg/db"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// 获取或创建频道
func GetOrCreateChannel(channelName string, userID uint, memberIDs []uint, isGroup bool) (models.Channel, error) {
	var channel models.Channel
	var members []uint

	// 群聊逻辑
	if isGroup {
		members = memberIDs // 直接使用传入的成员ID列表
		if channelName == "" {
			channelName = fmt.Sprintf("group_channel_%d", time.Now().Unix()) // 如果没有指定群聊名，生成一个
		}
	} else {
		// 私聊逻辑: 用两个用户的 ID 生成唯一频道名
		if len(memberIDs) != 1 {
			return models.Channel{}, errors.New("私聊只能包含一个目标用户")
		}
		targetUserID := memberIDs[0]
		if userID < targetUserID {
			channelName = fmt.Sprintf("private_%d_%d", userID, targetUserID)
		} else {
			channelName = fmt.Sprintf("private_%d_%d", targetUserID, userID)
		}
		members = []uint{userID, targetUserID} // 私聊成员是当前用户和目标用户
	}
	membersStr, _ := json.Marshal(members)
	// 查找现有频道
	err := db.DB.Where("name = ?", channelName).First(&channel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 频道不存在，创建新频道
			channel = models.Channel{
				ID:      fmt.Sprintf("%s_%d", channelName, len(members)), // 自定义生成频道ID
				Name:    channelName,
				IsGroup: isGroup,
				Members: string(membersStr),
			}
			if err := db.DB.Create(&channel).Error; err != nil {
				return models.Channel{}, err
			}
		} else {
			return models.Channel{}, err
		}
	}
	return channel, nil
}

// 保存聊天记录
func SaveMessage(channelID string, userID uint, username, content string) error {
	message := models.Message{
		ChannelID: channelID,
		UserID:    userID,
		Username:  username,
		Content:   content,
	}

	if err := db.DB.Create(&message).Error; err != nil {
		return err
	}
	return nil
}

// 获取频道消息
func GetChannelMessages(channelID string) ([]models.Message, error) {
	var messages []models.Message
	err := db.DB.Where("channel_id = ?", channelID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
