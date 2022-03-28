package main

import (
	"fmt"
	"strings"
)

func ProcessTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session Closed by close command")
		return false
	}else if strings.HasPrefix(str, "@shutdown"){
		fmt.Println("Session Shutdown")
		exitChan <- 0
		return false
	}

	fmt.Println("command process = ", str)
	return true
}
