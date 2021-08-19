package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	once := &sync.Once{}

	//开启两个线程，但是do中的方法只运行了一次
	go doSomething(once)
	go doSomething(once)

	time.Sleep(1e9)
}

func doSomething(once *sync.Once)  {	//传入的需要是一个指针，才能保证两个线程使用的是同一个once
	fmt.Println("Starting")
	once.Do(func() {
		fmt.Println("do something")
	})

	fmt.Println("ending")
}
