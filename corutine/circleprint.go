package main

import "fmt"

// main
/* @Description: 实现循环打印
*/
func main() {
	c := make(chan int)

	go PrintSingle(c)
	for i:=1; i<50;i++ {
		c<-i
		if i%2==0 {
			fmt.Println("main print = ", i)
		}
	}
	//以0结尾，告诉go func没有数据了，可以结束了
	c<-0
	x := <-c	//接收结束的数据
	fmt.Println("got go func return by x =", x)
}

func PrintSingle(c chan int)  {
	for {
		data := <-c
		if data	== 0{
			fmt.Println("get main ending, i will return")
			break
		}
		if data%2 == 1 {
			fmt.Println("go func print = ", data)
		}
	}
	c<-0
}
