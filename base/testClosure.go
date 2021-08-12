package main

import "fmt"

func main() {
	varClosure()
	surePartSafe()
	funcAsParam()
	funcAsReturn()
}

/**
* 定义匿名函数
* varClosure
* @Description:
*/
func varClosure()  {
	add := func(a,b int) int{
		return a+b
	}

	fmt.Println(add(1,2))
}

/**
* 闭包函数，保证局部变量的安全性
* surePartSafe
* @Description:
*/
func surePartSafe()  {
	j := 1
	f := func(){
		i := 1
		fmt.Println("i, j: ",i, j)		//闭包持有的j为外部变量的引用
	}

	f()
	j += 2
	f()
}

/**
* 函数作为入参，可以实现一个函数执行不同的算法，提升了代码的复用性
* funcAsParam
* @Description:
*/
func funcAsParam()  {
	add1 := func(a,b int) int {
		return a+b
	}

	add2 := func(a,b int) int{
		return a*2 + b
	}

	var a,b int = 1,2
	handleAdd(a,b, add1)
	handleAdd(a,b, add2)
}

/**
* 通过传入不同的函数，来实现不同的逻辑
* handleAdd
* @Description: 
* @param a
* @param b
* @param callFunc
*/
func handleAdd(a,b int, callFunc func(int, int) int)  {
	fmt.Println(callFunc(a,b))
}

/**
* 返回函数的延时执行
* funcAsReturn
* @Description:
*/
func funcAsReturn()  {
	add := funcRe(1,2)
	//.....可以有其他的逻辑

	fmt.Println(add())	//可实现延时执行
}

/**
* 返回一个函数
* funcRe
* @Description:
* @param a
* @param b
* @return func() int
*/
func funcRe(a,b int) func() int {
	return func() int {
		return a +b;
	}
}