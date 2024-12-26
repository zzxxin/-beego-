package services

import (
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQConn 封装RabbitMQ连接
var RabbitMQConn *amqp.Connection

// InitRabbitMQ 初始化RabbitMQ连接
func InitRabbitMQ() {
	var err error
	RabbitMQConn, err = amqp.Dial("amqp://myuser:mypass@127.0.0.1:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
}

// CloseRabbitMQ 关闭连接
func CloseRabbitMQ() {
	if RabbitMQConn != nil {
		RabbitMQConn.Close()
	}
}

// SendMessage 发送Mq消息 queueName 队列名称   message 消息内容
func SendMessage(queueName, message string) {

	if RabbitMQConn == nil || RabbitMQConn.IsClosed() {
		log.Fatalf("RabbitMQ connection is not open")
	}

	ch, err := RabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // 无等待
		nil,       // 额外参数
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf(" 成功发送消息 %s", message)
}

// ReceiveMessages 从指定的队列接收 MQ 消息，并通过通道返回
func ReceiveMessages(queueName string) (<-chan string, error) {
	// 创建一个通道来返回消息
	messageChan := make(chan string)

	// 打开 RabbitMQ 的 channel
	ch, err := RabbitMQConn.Channel()
	if err != nil {
		return nil, err
	}

	// 申明队列（确认队列已存在）
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // 无等待
		nil,       // 额外参数
	)
	if err != nil {
		return nil, err
	}

	// 注册消费者
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者
		true,   // 自动应答
		false,  // 独占
		false,  // 无等待
		false,  // 额外参数
		nil,    // 额外参数
	)
	if err != nil {
		return nil, err
	}

	// 使用 goroutine 来监听消息队列
	go func() {
		defer ch.Close() // 保证退出时关闭 channel
		for d := range msgs {
			messageChan <- string(d.Body) // 将消息内容发送到 messageChan 中
		}
		close(messageChan) // 当消息队列关闭时关闭通道
	}()

	return messageChan, nil
}
