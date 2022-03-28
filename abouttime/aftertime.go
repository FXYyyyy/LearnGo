package main

import (
	"fmt"
	"time"
)

func main() {
	exit := make(chan int)

	fmt.Println("start...")

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("there is 3s past")
		fmt.Println("I will tell main to exit")
		exit<-0
	})

	<-exit
	fmt.Println("Get exit msg, exit Now!")
}
