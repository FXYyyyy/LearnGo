package mq

import (
	_const "firstGo/testrabbitmq/const"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// RabbitMQ
/*  @Description:rabbitMQ结构体
*/
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

// NewRabbitMQ
/* @Description: 创建结构体实例
*  @param queueName
*  @param exchange
*  @param key
*  @return *RabbitMQ
*/
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: _const.MQURL}
}

// Destroy
/* @Description: 断开channel 和 connection
*  @receiver r
*/
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// FailOnError 
/* @Description: 对错误进行日志的打印
*  @param err
*  @param errMsg
*  @param msg
 */
func FailOnError(err error, errMsg string, msg string) {
	if err != nil {
		log.Fatalln("%s: %s", errMsg, err)
	} else {
		fmt.Println("Success: ", msg)
	}
}

