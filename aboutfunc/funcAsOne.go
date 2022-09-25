package main

import "fmt"

//定义一个调用器的接口
type Invoker interface {
	Call(interface{})
}

//继承调用器
type MyStruct struct {
}
func (m *MyStruct)Call(p interface{})  {
	fmt.Println("From MyStruct : ", p)
}

//定义了一个函数类型
type FuncCaller func(interface{})
func (f FuncCaller)Call(p interface{})  {
	f(p)	//调用FuncCaller函数本体
}

func main() {
	var invoker Invoker

	s := new(MyStruct)
	invoker = s
	invoker.Call("hello")

	//将匿名函数转化成FuncCaller类型，再赋值给invoker
	invoker = FuncCaller(func(p interface{}) {
		fmt.Println("from funcCaller = ", p)
	})
	invoker.Call("zhizhi")
}
