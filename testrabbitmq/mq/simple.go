package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// NewRabbitMQSimple
/* @Description: 创建简单模式下RabbitMQ实例
*  @param queueName
*  @return *RabbitMQ
 */
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")

	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	FailOnError(err, "failed to connect rabbitmq!", "Success to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	FailOnError(err, "failed to open a channel", "Success to open a channel")
	return rabbitmq
}

// QueueDeclareSimple
/* @Description: 申请队列，如果存在则跳过
*  @receiver r
 */
func (r *RabbitMQ)QueueDeclareSimple()  {
	//申请队列，如果存在则跳过
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false,		//是否持久化
		false,	//是否自动删除
		false,		//是否j具有排他性
		false,		//是否阻塞处理
		nil)			//其他属性
	FailOnError(err, "Queue Declare Failed", "Queue Declare Success")
}


// PublishSimple
/* @Description: Simple模式的队列生产者
*  @receiver r
*  @param message
 */
func (r *RabbitMQ)PublishSimple(message string)  {
	//申请队列，如果存在则跳过
	r.QueueDeclareSimple()

	//发送消息到队列中
	err := r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
	FailOnError(err, "message publish Failed", "message publish Success")
}


// ConsumeSimple
/* @Description: Simple模式的消费者
*  @receiver r
 */
func (r *RabbitMQ)ConsumeSimple()  {
	//申请队列，如果存在则跳过
	r.QueueDeclareSimple()

	//接收消息
	msg, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil)
	FailOnError(err, "get msg failed", "get msg Success")

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)
			fmt.Println("标准输出: ", string(d.Body))
		}
	}()

	fmt.Println("[*] Wait for message, To exit press ctrl + c")
	<-forever
}
