package main

import (
	"fmt"
)

/**
* 测试简单的chan
* main
* @Description:
*/
func main() {
	//testChan()
	testBufChan()
}

/**
* 测试一个简单的channem
* testChan
* @Description: 
*/
func testChan()  {
	chs := make([]chan int, 10)
	for i:=0; i<10; i++ {
		chs[i] = make(chan int)
		go addWithChan(1, i, chs[i])	//开启子线程
	}

	for k,ch := range chs {
		x := <- ch
		fmt.Println(k, ch, x)
	}

	fmt.Println("End")
}
func addWithChan(a,b int, ch chan int)  {
	c := a + b
	fmt.Println("ret give chan : ", c)
	ch <- c
	close(ch)
}


//=======================================
/**
* 测试一个带缓冲的通道
* testBufChan
* @Description:
*/
func testBufChan()  {
	ch := make(chan int, 20)
	go addIntToChan(ch)

	for c :=range ch {
		println(c)
	}
}
func addIntToChan(ch chan int)  {
	for i:=0; i<100;i++ {
		ch<-i
	}
	close(ch)	//需显示关闭通道，都则主程序不知道子协程什么时候执行完，从一个空通道接受数据，会报错死锁
}