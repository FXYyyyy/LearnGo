package main

import (
	"fmt"
	"net"
)

func telnet_server(address string, exitchan chan int) {
	//根据给定地址进行监听
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Listen the address = ", address, "failed! err = ", err)
		exitchan <- 1
	}

	fmt.Println("Listen Success, address = ", address)
	defer listen.Close()

	for  {
		//新链接没有到来的时候，accept是阻塞的
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Got the connect Failed, err = ", err)
			continue
		}

		go HandleSession(conn, exitchan)
	}
}
