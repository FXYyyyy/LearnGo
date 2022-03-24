package main

import (
	"fmt"
	"time"
)

/**
* 简单的开启协程测试
* main
* @Description:
*/
func main() {
	fmt.Println("Start main")
	go goAdd(1, 2)	//开启一个用户线程，引入了不确定性，主函数无法捕捉返回
	time.Sleep(1e9)

	fmt.Println("End main")

	go running()
	//此处阻塞了main，所以running会一直输出
	var input string
	fmt.Scanln(&input)
	fmt.Println("justInput = ", input)

}

func goAdd(a,b int){
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick tok", times)

		time.Sleep(time.Second)
	}
}
