package main

import (
	"fmt"
	"time"
)

func main() {
	//一个打点器，每500ms会向c chan中发送一次时间
	ticker := time.NewTicker(500*time.Millisecond)

	//一个计时器，在2s后向c chan发送一次时间
	stopTimer := time.NewTimer(2*time.Second)

	i := 0

	for {
		select {
		case x := <-ticker.C :	//监听打点器的C chan
			fmt.Println("get ticker = ",x)
			i++
		case s := <-stopTimer.C:	//监听计时器的C chan
			fmt.Println("stop timer is compete with s =", s)
			goto breakHere	//一个代码块的跳转
			/*	代码停止不了，说明timer的C只会有一次的通知
			if i == 6 {
				goto breakHere	//一个代码块的跳转
			}
			 */
		}
	}

	breakHere:
		fmt.Println("main stop with i = ", i)
}
