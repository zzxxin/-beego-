package socket

import (
	"beegoweb/controllers"
	"beegoweb/models"
	services "beegoweb/service"
	"fmt"
	"strconv"
)

type ChatController struct {
	controllers.BaseController
}

// 渲染聊天室页面
func (c *ChatController) ChatRoom() {
	var operator uint
	if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
		operator = op
	}

	// 获取当前用户信息
	userInfo, _ := models.GetUserByID(operator)

	c.Data["user_info"] = userInfo
	c.TplName = "socket/chatroom.tpl"
}

// GetChannel 获取或创建频道
func (c *ChatController) GetChannel() {
	var operator uint
	if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
		operator = op
	}
	// 获取必要的参数
	isGroup, _ := c.GetBool("is_group", false)
	channelName := c.GetString("channel_name", "")
	memberIDs := c.GetStrings("member_ids[]") // 成员ID列表，群聊时为多个，私聊时为1个
	userID := operator                        // 当前用户ID
	fmt.Printf("收到的用户ID是：%v\n", memberIDs)

	// 处理成员ID列表
	members := make([]uint, len(memberIDs))
	for i, id := range memberIDs {
		memberID, _ := strconv.Atoi(id)
		members[i] = uint(memberID)
	}

	// 调用创建或获取频道的方法
	channel, err := services.GetOrCreateChannel(channelName, userID, members, isGroup)
	if err != nil {
		fmt.Printf("错误信息为%e", err)
		c.Data["json"] = map[string]string{"error": "频道创建失败"}
		c.ServeJSON()
		return
	}

	// 获取该频道的聊天记录
	messages, _ := services.GetChannelMessages(channel.ID)

	c.Data["json"] = map[string]interface{}{
		"channel":  channel,
		"messages": messages,
	}
	c.ServeJSON()
}

// GetMessages 获取聊天记录
func (c *ChatController) GetMessages() {
	channelID := c.GetString("channel_id")
	messages, err := services.GetChannelMessages(channelID)
	if err != nil {
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		// 将时间格式化为字符串
		formattedMessages := make([]map[string]interface{}, len(messages))
		for i, msg := range messages {
			formattedMessages[i] = map[string]interface{}{
				"id":         msg.ID,
				"channel_id": msg.ChannelID,
				"user_id":    msg.UserID,
				"username":   msg.Username,
				"content":    msg.Content,
				"timestamp":  msg.Timestamp.Format("2006-01-02 15:04:05"), // 格式化时间
			}
		}
		c.Data["json"] = formattedMessages
	}
	c.ServeJSON()
}
