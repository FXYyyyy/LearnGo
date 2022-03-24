package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg sync.WaitGroup
)

func incCount()  {
	defer wg.Done()

	for i:=0; i<2; i++ {
		value := count

		//goroutine暂停，退出执行队列，让给其他等待的goroutine，为了模拟资源竞争
		runtime.Gosched()

		value++
		count = value
	}
}

func incCountByAtomic() {
	defer wg.Done()
	for i:=0; i<2; i++ {
		//原子操作，安全的对count+2
		atomic.AddInt32(&count, 2)
		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)
	//go incCount()
	//go incCount()
	go incCountByAtomic()
	go incCountByAtomic()
	wg.Wait()
	//运行结果不稳定，可能是2、3、4
	fmt.Println("now count = ", count)
}
