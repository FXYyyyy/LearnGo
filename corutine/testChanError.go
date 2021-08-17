package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	chanTimeOutHandle()

	chanCloseHandle()
}

/**
* 处理超时机制的实现，防止某个通道写入超时，等待太久，导致崩溃
* chanTimeOutHandle
* @Description:
*/
func chanTimeOutHandle()  {
	ch := make(chan int, 1)

	timeOut := make(chan bool, 1)

	go func() {
		time.Sleep(1e9)	//睡眠1s
		timeOut<-true
	}()

	select {
	case <-ch:
		fmt.Println("got ch")
		case timeOut<-true :
			fmt.Println("timeout 1s, exit waiting ch")
	default:
		fmt.Println("default")
	}
}

/**
* 对某个通道是否已经关闭做校验
* chanCloseHandle
* @Description:
*/
func chanCloseHandle()  {
	ch := make(chan int, 1)

	//发送方
	go func() {
		for i:=0; i<5; i++{
			fmt.Println("send = ", i)
			ch<-i
		}
		close(ch)
		fmt.Println("send close")
	}()

	//接收方
	for {
		num, ok := <-ch
		if !ok{
			fmt.Println("got close")
			break
		}

		fmt.Println("got num = ", num)
	}

	fmt.Println("cpu num = ", runtime.NumCPU())	//获取cpu核心数
}
