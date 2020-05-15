package main

import (
	"fmt"
	"lph-rabbitmq/RabbitMQ"
)

func main() {

	//订阅模式
	//rabbitmq := RabbitMQ.NewRabbitMQPubSub("lphExchange","")
	//fanout订阅模式
	//rabbitmq.PublishPub("Hello","fanout")
	//路由模式
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("lphExchange2","one")
	rabbitmq2 := RabbitMQ.NewRabbitMQPubSub("lphExchange2","two")
	//direct路由模式
	rabbitmq.PublishPub("Hello","direct")
	rabbitmq2.PublishPub("Hello2","direct")
	fmt.Println("sub发送成功")
}
