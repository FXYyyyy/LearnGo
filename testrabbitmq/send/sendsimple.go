package main

import (
	_const "firstGo/testrabbitmq/const"
	"firstGo/testrabbitmq/mq"
	"strconv"
	"time"
)

func main()  {
	r := mq.NewRabbitMQSimple(_const.TestQueueName)

	for i:=0; i<100; i++ {
		r.PublishSimple("lili 是猪" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}

	r.Destroy()
}