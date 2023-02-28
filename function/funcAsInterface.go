package main

import "fmt"

// main
/*
 @Description: 函数类型实现接口
*/
func main() {
	var invoker Invoker

	z := new(Zhi)

	invoker = z

	invoker.Call("hello")

	//将匿名函数转换成FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("From function: ", v)
	})

	invoker.Call("hello")
}

type Invoker interface {
	Call(interface{})
}

type Zhi struct {
}

func (z *Zhi) Call(p interface{}) {
	fmt.Println("from zhizhi", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}
