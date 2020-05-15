package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//根据自己的rabbitmq 修改
const MQURL = "amqp://admin:123456@192.168.0.22:5672"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	MqUrl string
}

//创建RabbitMQ实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Key: key, Exchange: exchange, MqUrl: MQURL}
	var err error

	//创建
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "connection error")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "channel error")
	return rabbitmq
}

//断开
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//简单模式
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

//订阅模式
func NewRabbitMQPubSub(exchangeName string,routingKey string) *RabbitMQ {
	return NewRabbitMQ("",exchangeName,routingKey)
}

//检查队列
func (r *RabbitMQ) CheckQueue() {
	//申请队列，无则自动创建，有则跳过
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Print(err)
	}
}
//检查交换机
func (r *RabbitMQ) CheckExchange(kind string) {
	//申请交换机，无则自动创建，有则跳过
	 err := r.channel.ExchangeDeclare(
		r.Exchange,
		//类型广播类型
		//"fanout",
		 kind,

		true,

		false,

		false,

		false,

		nil,
	)
	if err != nil {
		fmt.Print(err)
	}
}

//简单模式生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	//申请队列，无则自动创建，有则跳过
	r.CheckQueue()
	//发送消息到队列中
	r.Publish(message,r.QueueName)

}

//消费者
func (r *RabbitMQ) Consume(name string)  {
	//接受消息
	msgs,err := r.channel.Consume(
		name,
		//用来区分消费者
		"",
		//是否自动应答
		true,
		//排他
		false,
		//如果这只为true 表示不能江同个connection中发送到消息传递给这个connection中的消费者
		false,
		//是否阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Print(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//增加逻辑代码
			log.Printf("Received a message:%s",d.Body)
			//fmt.Println(d.Body)
		}
	}()
	log.Printf("waiting for messages")
	<-forever
}

//简单模式消费者
func (r *RabbitMQ) ConsumeSimple()  {
	//申请队列，无则自动创建，有则跳过
	r.CheckQueue()
	//接受消息
	r.Consume(r.QueueName)
}

//订阅模式消费者
func (r *RabbitMQ) ConsumeSub(kind string)  {
	//申请队列，无则自动创建，有则跳过
	r.CheckExchange(kind)
	//接受消息
	q,err := r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)
	if err != nil {
		fmt.Print(err)
	}

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
		)
	//消费消息
	r.Consume(q.Name)
}

//订阅模式生产代码
func (r *RabbitMQ) PublishPub(message string,kind string) {
	//申请交换机，无则自动创建，有则跳过
	r.CheckExchange(kind)
	//发送消息到队列中
	r.Publish(message,r.Key)
}

//通用生产代码
func (r *RabbitMQ) Publish(message string,key string) {
	r.channel.Publish(
		r.Exchange,
		//r.QueueName,
		key,
		//如果为true 会根据exchange 类型 和 routkey 规则，如果找不到符合的 则返回给发送者
		false,
		//如果为true 如果无消费者 则返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
}

