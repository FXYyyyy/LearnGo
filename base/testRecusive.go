package main

import (
	"fmt"
	"time"
)

func main() {
	n := 20

	decorator := spendTime1(fibonacci)
	fmt.Println(decorator(n))

	newFDec := spendTime1(newFibonacciByTail)
	fmt.Println(newFDec(n))
}

/**
* 原始计算斐波拉契
* fibonacci
* @Description:
* @param n
* @return int
*/
func fibonacci(n int) int {
	if n==1{
		return 0
	}
	if n==2{
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

type fiFuc func(int)int

/**
* 装饰器，计算耗时
* spendTime1
* @Description:
* @param f
* @return fiFuc
*/
func spendTime1(f fiFuc) fiFuc {
	return func(n int) int {
		startTime := time.Now()
		ret := f(n)
		fmt.Println("---spendTime =", time.Since(startTime))
		return ret
	}
}

/**
* 尾递归优化，记录原来通过递归调用计算结果转化为通过外部传递参数初始化，再传递给下次尾递归调用不断累加
* newFibonacciByTail
* @Description:
* @param n
* @return int
*/
func newFibonacciByTail(n int) int {
	return fibonacciTail(n, 0, 1)
}
func fibonacciTail(n, first, second int) int {
	if n<2{
		return first
	}

	return fibonacciTail(n-1, second, first+second)
}

