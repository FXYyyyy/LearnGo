package main

import (
	"fmt"
	"sync"
)

func main() {
	//simplePool()
	testPoolWithGroup()
}

/**
* simplePool
* @Description: 一个简单的临时对象池的使用
*/
func simplePool()  {
	pool := sync.Pool{
		New: func() interface{} {	//如果pool临时对象池里没有数据，则通过New方法，创建一个数据返回，保证始终有返回
			return "Hell, Pool"
		},
	}

	value := "zhizhi"	//
	pool.Put(value)

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}

/**
* shareString
* @Description: 利用pool，实现协程间的对象资源共享
* @param pool
* @param deferFunc
*/
func shareString(i string, pool *sync.Pool, deferFunc func())  {
	defer deferFunc()
	str := "zhihzi-" + i
	pool.Put(str)	//将str存入pool，实现该协程的对象，与其他协程的共享
}
func testPoolWithGroup()  {
	var wg sync.WaitGroup
	wg.Add(2)

	pool := sync.Pool{
		New: func() interface{} {
			return "aaa"
		},
	}

	go shareString("1", &pool, wg.Done)
	go shareString("2", &pool, wg.Done)
	wg.Wait()

	fmt.Println(pool.Get())	//获取子协程放入的对象
	fmt.Println(pool.Get())	//获取子协程放入的对象
}