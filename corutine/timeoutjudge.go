package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	other := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch :
				fmt.Println("num = ", num)
			case <-time.After(10*time.Second):	//10s超时
				fmt.Println("10 second over, there is timeout")
				quit <- true	//向quit发送信号
			case x := <-other :
				fmt.Println("other also get the num = ", x)
			}
			fmt.Println("end once circle")
		}
	}()

	for i:=0;i<5;i++ {
		ch <- i
		other <- i
	}

	flag := <-quit
	fmt.Println("there is a signal : ", flag, ", the main ending")
}
