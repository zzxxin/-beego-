package api

import (
	"beegoweb/controllers"
	services "beegoweb/service"
	"beegoweb/utils"
	"fmt"
	"log"
)

type TestController struct {
	controllers.BaseController
}

// 测试接口类型数据返回
func (c *TestController) TestApi() {

	//  测试发送mq消息内容
	services.SendMessage("test", "test Mq message")

	utils.JSONResponse(&c.Controller, 401, "用户未登录", nil)
}

// 消费mq
func (c *TestController) QueueMq() {
	messageChan, err := services.ReceiveMessages("test")
	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}

	// 从通道中读取消息并处理
	for msg := range messageChan {
		fmt.Printf("拿到的消息是: %s\n", msg)
	}

	log.Println("No more messages, exiting...")
}
