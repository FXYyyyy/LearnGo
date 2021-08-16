package main

import (
	"errors"
	"fmt"
)

func main() {
	//testError()
	//testDefer()
	testRecover();
	fmt.Println("exit in main")
}

/**
* 测试错误
* testError
* @Description:
*/
func testError()  {
	x := -1
	y := 2
	z, err := addWithErr(x, y)
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("a + b = ", z)
	}
}

func addWithErr(a, b int) (n int, err error) {
	if(a < 0 || b < 0){
		err = errors.New("不支持负数的操作")
		return
	}

	n = a+b
	return
}

/**
* 测试defer修饰的兜底逻辑
* testDefer
* @Description:
*/
func testDefer()  {
	defer func() {fmt.Println("last run")}()

	defer fmt.Println("执行完了除法")
	i := 1
	j := 0
	ret := i/j

	defer fmt.Println("ret = ", ret)
}

/**
* 测试recover捕捉
* testRecover
* @Description:
*/
func testRecover()  {
	defer func() {
		fmt.Println("resolve panic")
		if err := recover(); err != nil{
			fmt.Printf("get a panic: %v\n", err)
		}
	}()

	i := 1
	j := 0
	ret := i/j
	fmt.Println("i/j = ", ret)
}