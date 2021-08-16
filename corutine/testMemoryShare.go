package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var Counter int = 0	//共享内存，写入、读取时需要加锁

/**
* 内存共享小测试
* main
* @Description:
*/

func main() {
	startTime := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i<10; i++ {
		addWithLock(1, i, lock)
	}

	for {
		lock.Lock()
		c := Counter
		lock.Unlock()
		println(c)
		runtime.Gosched()	//用于让cpu让出时间片
		if c >= 10 {
			break
		}
	}

	endTime := time.Now()
	consume := endTime.Sub(startTime).Seconds()
	fmt.Println("耗时：", consume)
}

func addWithLock(a, b int, lock *sync.Mutex)  {
	c := a + b
	lock.Lock()
	Counter++
	lock.Unlock()
	fmt.Printf("%d + %d = %d\n", a, b, c)
}
