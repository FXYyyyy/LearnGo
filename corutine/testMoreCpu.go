package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
* 利用多核cpu实现并行计算
* main
* @Description:
*/
func main() {
	start := time.Now()
	cpuNum := runtime.NumCPU()	//获取cpu的核心数
	runtime.GOMAXPROCS(1)	//设置运行时，可用的cpu的最大核心数

	chs := make([]chan int, cpuNum)
	for i:=0; i<cpuNum; i++{
		chs[i] = make(chan int, 1)
		go setCh(i, chs[i])
	}

	for _,ch := range chs{
		fmt.Println("got i = ", <-ch)
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("END Main with consume = ", consume)	//通过耗时可以看出cpu核心数的不同对代码效率的影响
}

func setCh(i int, ch chan int)  {
	defer close(ch)

	fmt.Println("send i = ", i)
	ch<-i
}
