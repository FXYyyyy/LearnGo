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
}

func goAdd(a,b int){
	fmt.Printf("%d + %d = %d", a, b, a+b)
}
