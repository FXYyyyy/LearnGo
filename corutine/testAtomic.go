package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg1 sync.WaitGroup
)

func main() {
	wg1.Add(2)

	go dowork("lili")
	go dowork("zhizhi")

	time.Sleep(1* time.Second)
	fmt.Println("main shut Now")
	//安全写
	atomic.StoreInt64(&shutdown, 1)
	wg1.Wait()
}

func dowork(name string)  {
	defer wg1.Done()
	for {
		fmt.Println("doing work = ", name)
		time.Sleep(250*time.Millisecond)

		//安全读
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("shuting down = ", name)
			break
		}
	}
}
