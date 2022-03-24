package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg2 sync.WaitGroup
	lock sync.Mutex
)

// main
/* @Description: 互斥锁
 */
func main() {
	wg2.Add(2)

	go MutexTest(1)
	go MutexTest(2)

	wg2.Wait()
	fmt.Println("counter = ", counter)
}

func MutexTest(id int)  {
	defer wg2.Done()
	for i:=0;i<id;i++ {
		lock.Lock()
		value := i
		runtime.Gosched()
		value++
		counter = value
		lock.Unlock()
	}
}
