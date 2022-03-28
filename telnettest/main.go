package main

import "os"

func main() {
	exitchan := make(chan int)

	go telnet_server("127.0.0.1:7001", exitchan)

	code := <-exitchan

	//标记程序返回值并退出zhi
	os.Exit(code)
}
