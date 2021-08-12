package main

import (
	"fmt"
	"time"
)

func main() {
	a, b := 1,2

	//spend装饰器，添加了对函数耗时的计算，只需要将不同的函数传入即可，而不需要改动计算耗时的逻辑
	decoratorAdd := spendTime(add)	//装饰器模式，在add函数上加其他的逻辑，而不会影响这个函数本身的计算
	ret := decoratorAdd(a, b)
	fmt.Println(ret)

	decoratorMulti := spendTime(multi)
	fmt.Println(decoratorMulti(a,b))

	decoratorBite := spendTime(bite)
	fmt.Println(decoratorBite(a, b))
}

/**
* 两数想加的函数
* add
* @Description:
* @param a
* @param b
* @return int
*/
func add(a,b int) int {
	return a+b
}

func multi(a,b int) int {
	return a*b
}

func bite(a, b int) int {
	return a<<b
}

type twoIntRet func(int, int)int	//对该类型的函数进行重命名的定义

/**
* 装饰器模式，求函数的耗时
* spendTime
* @Description:
* @param f
* @return addFunc
*/
func spendTime(f twoIntRet) twoIntRet {
	return func(a int, b int) int {
		startTime := time.Now()
		ret := f(a, b)
		fmt.Println("spend Time = ", time.Since(startTime))
		return ret
	}
}
