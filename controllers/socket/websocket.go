package socket

import (
	"beegoweb/models"
	services "beegoweb/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	web.Controller
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	broadcast   = make(chan Message)     // 广播通道
	mu          sync.Mutex               // 锁，确保并发安全
	clients     = make(map[*Client]bool) // 已连接的客户端
	onlineUsers = make(map[uint]bool)    // 在线用户
)

type Message struct {
	UserId       uint   `json:"userId"`       // 发送消息的用户ID
	Username     string `json:"username"`     // 用户名
	Message      string `json:"message"`      // 消息内容
	Type         string `json:"type"`         // 消息类型
	TargetUserId uint   `json:"targetUserId"` // 目标用户ID（私聊时使用）
	ChannelId    string `json:"channel_id"`   // 频道ID，私聊或群聊的标识
}

type Client struct {
	Conn      *websocket.Conn // WebSocket连接
	UserId    uint            // 客户端的登录用户ID
	ChannelId string          // 当前用户加入的频道ID（默认为公共频道）
}

// Get 处理 WebSocket 连接
func (c *WebSocketController) Get() {
	// 获取当前登录用户的ID
	var operator uint
	if op, ok := c.Ctx.Input.GetData("userId").(uint); ok {
		operator = op
	} else {
		log.Println("用户未登录，拒绝连接")
		return
	}

	// 获取当前用户信息
	nowUserInfo, _ := models.GetUserByID(operator)

	// 升级为 WebSocket
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Printf("WebSocket 升级失败: %v", err)
		return
	}
	defer ws.Close()

	// 初始化链接
	client := &Client{Conn: ws, UserId: operator, ChannelId: ""} // 默认加入公共频道

	// 用户上线，标记为在线状态
	mu.Lock()
	clients[client] = true
	onlineUsers[operator] = true
	mu.Unlock()

	// 发送上线通知
	broadcast <- Message{
		UserId:   operator,
		Username: nowUserInfo.UserName,
		Type:     "system",
		Message:  "用户 " + nowUserInfo.UserName + " 上线了",
	}

	// 发送在线用户列表
	sendOnlineUsers()

	defer func() {
		// 用户下线，移除客户端
		mu.Lock()
		delete(clients, client)
		delete(onlineUsers, operator)
		mu.Unlock()

		// 发送下线通知
		broadcast <- Message{
			UserId:   operator,
			Username: nowUserInfo.UserName,
			Type:     "system",
			Message:  "用户 " + nowUserInfo.UserName + " 下线了",
		}
		// 更新在线用户列表
		sendOnlineUsers()
	}()

	// 不断监听从客户端发送的消息
	for {
		var msg map[string]interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("读取消息失败: %v", err)
			break
		}

		// 处理频道切换
		if msg["type"] == "switch_channel" {
			channelID, ok := msg["channel_id"].(string)
			if ok {
				mu.Lock()
				client.ChannelId = channelID // 更新客户端的频道ID
				mu.Unlock()
				log.Printf("用户 %d 切换到频道 %s", client.UserId, channelID)
				continue
			}
		}

		// 处理普通消息
		userIdStr, ok := msg["userId"].(string)
		if !ok {
			log.Printf("userId 不是字符串类型")
			continue
		}
		userIdUint, err := strconv.ParseUint(userIdStr, 10, 32)
		if err != nil {
			log.Printf("无法转换 userId: %v", err)
			continue
		}

		targetUserIdStr, ok := msg["targetUserId"].(string)
		var targetUserIdUint uint64
		if ok {
			targetUserIdUint, _ = strconv.ParseUint(targetUserIdStr, 10, 32)
		}

		// 创建消息
		message := Message{
			UserId:       uint(userIdUint),
			Username:     msg["username"].(string),
			Message:      msg["message"].(string),
			Type:         "message",
			TargetUserId: uint(targetUserIdUint),
			ChannelId:    client.ChannelId, // 使用当前客户端的频道ID
		}

		// 发送消息到广播通道
		broadcast <- message
	}
}

// 广播消息给所有客户端
func handleMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		if msg.Type == "message" {
			// 保存消息
			err := services.SaveMessage(msg.ChannelId, msg.UserId, msg.Username, msg.Message)
			if err != nil {
				log.Printf("保存消息失败: %v", err)
				continue
			}
		}
		for client := range clients {
			// 检查是否广播给自己
			if client.UserId == msg.UserId {
				continue // 跳过发送给自己
			}
			if msg.TargetUserId > 0 && client.UserId == msg.TargetUserId && client.ChannelId == msg.ChannelId {
				fmt.Printf("走的私聊发送\n")
				// 私聊消息，只发给目标用户
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("发送私聊消息失败: %v", err)
					client.Conn.Close()
					delete(clients, client)
				}
			} else if msg.TargetUserId == 0 && client.ChannelId == msg.ChannelId {
				fmt.Printf("走的群聊发送\n")
				// 群聊消息，发给同一个频道的所有用户
				err := client.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("发送群聊消息失败: %v", err)
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
		mu.Unlock()
	}
}

func init() {
	// 启动消息处理 goroutine
	go handleMessages()
}

// 发送在线用户列表给所有客户端
func sendOnlineUsers() {
	mu.Lock()
	var userIDs []uint
	for userID := range onlineUsers {
		userIDs = append(userIDs, userID)
	}
	allUserList, _ := models.GetUserAllByQuery("1=1")
	userList := map[string]interface{}{
		"user_list": allUserList,
		"bind_ids":  userIDs,
	}
	mu.Unlock()

	userListJSON, _ := json.Marshal(userList)
	broadcast <- Message{
		Type:     "online_users",
		UserId:   0,
		Username: "系统通知",
		Message:  string(userListJSON),
	}
}
