package main

import "lph-rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("lph")
	rabbitmq.ConsumeSimple()
}
