package main

import (
	"fmt"
	"lph-rabbitmq/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("lph")
	rabbitmq.PublishSimple("Hello")
	fmt.Println("发送成功")
}
