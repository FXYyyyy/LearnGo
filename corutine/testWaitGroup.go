package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	f := func() {
		wg.Done()
	}
	for i:=0; i<10; i++ {
		go AddWithDone(i, f)
	}

	wg.Wait()

	fmt.Println("ENd Main")
}

func AddWithDone(i int, f func())  {
	defer f()
	fmt.Println("cur = ", i)
}
