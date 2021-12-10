package mq

import (
	_const "firstGo/testrabbitmq/const"
	"fmt"
	"github.com/streadway/amqp"
)

// NewRabbitMqPubSub
/* @Description: 订阅模式下创建一个rabbit实例
*  @param exChangeName
*  @return *RabbitMQ
*/
func NewRabbitMqPubSub(exChangeName string) *RabbitMQ {
	r := NewRabbitMQ("", exChangeName, "")

	var err error
	r.conn, err = amqp.Dial(_const.MQURL)
	FailOnError(err, "failed to connect rabbitmq!", "Success to connect rabbitmq!")

	r.channel, err = r.conn.Channel()
	FailOnError(err, "failed to open a channel", "success to open a channel")

	return r
}

// ExChangeDeclarePub
/* @Description: 尝试建立交换机
*  @receiver r
*/
func (r *RabbitMQ)ExChangeDeclarePub()  {
	//尝试建立交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",	//广播
		true,
		false,
		false,
		false,
		nil)
	FailOnError(err, "Failed to Declare ExChange", "Success to Declare ExChange")
}

// PublishPub
/* @Description: 订阅模式下发送消息
*  @receiver r
*  @param message
*/
func (r *RabbitMQ)PublishPub(message string)  {
	//建立交换机
	r.ExChangeDeclarePub()

	//发送消息
	err := r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
	FailOnError(err, "PubSub Send msg Failed", "PubSub Send msg Success")
}

// ConsumePub
/* @Description: 订阅模式的消费端
*  @receiver r
*/
func (r *RabbitMQ)ConsumePub() {
	//尝试建立交换机
	r.ExChangeDeclarePub()

	//创建队列
	q, err := r.channel.QueueDeclare(
		"",	//随机产生名字
		false,
		false,
		true,
		false,
		nil)
	FailOnError(err, "Pub Declare queue Failed", "Pub Declare queue Success")

	//绑定队列至exchange
	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil)
	FailOnError(err, "Bind Queue Failed", "Bind Queue Success")

	//消费
	msg, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	FailOnError(err, "Failed to consume msg", "Success to consume msg")

	//消息对应的操作
	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Println("get msg :" , string(d.Body))
		}
	}()
	fmt.Println("wait for msg, To exit press ctrl+c")
	<-forever
}
