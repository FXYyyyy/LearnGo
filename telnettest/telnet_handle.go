package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleSession(conn net.Conn, exitChan chan int)  {
	fmt.Println("Session start! ")

	//创建一个网络数据读取器
	reader := bufio.NewReader(conn)

	//接收数据
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read the string err = ", err)
			conn.Close()
			break
		}
		//去除尾部的回车
		str = strings.TrimSpace(str)
		if !ProcessTelnetCommand(str, exitChan) {
			conn.Close()
			break
		}

		//在客户端连接返回数据
		conn.Write([]byte(str + "ohuo" + "\r\n"))
	}

	fmt.Println("Session End!")
}
