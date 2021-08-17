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
	/*
	ch := make(chan int, 20)
	//go addIntToChan(ch)
	go addIntToSingleTrack(ch)
	 */
	ch := returnSingleChan()

	for c :=range ch {	//通道不关闭，此处会陷入死循环
		println(c)
	}
}
func addIntToChan(ch chan int)  {	
	for i:=0; i<30;i++ {
		ch<-i
	}
	close(ch)	//需显示关闭通道，都则主程序不知道子协程什么时候执行完，从一个空通道接受数据，会报错死锁
}
func addIntToSingleTrack(ch chan<- int)  {	//单向通道，该函数只能对该通道纪念性写操作
	for i:=0; i<30;i++ {
		ch<-i
	}
	close(ch)
}

func returnSingleChan() <-chan int {	//该函数只负责通道的写，返回读的通道，使函数指责清晰
	ch := make(chan int, 30)	//此处的缓存，如果小于需放进去的变量数量，则会引起死锁
	for i:=0; i<30; i++ {
		ch<-i
	}
	close(ch)
	return ch
}
