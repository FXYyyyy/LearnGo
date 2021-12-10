package main

import (
	_const "firstGo/testrabbitmq/const"
	"firstGo/testrabbitmq/mq"
)

func main()  {
	r := mq.NewRabbitMQSimple(_const.TestQueueName)

	r.ConsumeSimple()
}
