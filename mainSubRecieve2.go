package main

import "lph-rabbitmq/RabbitMQ"

func main() {
	//订阅模式
	//rabbitmq := RabbitMQ.NewRabbitMQPubSub("lphExchange","")
	//rabbitmq.ConsumeSub("fanout")
	//路由模式
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("lphExchange2","two")
	rabbitmq.ConsumeSub("direct")
}
