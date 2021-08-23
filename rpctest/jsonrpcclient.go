package main

import (
	."firstGo/rpctest/utils"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	conn,err := net.DialTimeout("tcp", "localhost:8080", 30*time.Second)
	if err != nil {
		log.Fatalf("客户端发起连接失败：%v", err)
	}
	defer conn.Close()

	client := jsonrpc.NewClient(conn)

	var item Item
	client.Call("ServiceHandler.GetName", 1, &item)
	log.Printf("ServiceHandler.GetName 返回结果：%v\n", item)

	var response Response
	item = Item{2, "enen"}
	client.Call("ServiceHandler.SaveName", item, &response)
	log.Printf("ServiceHandler.SaveName 返回结果：%v\n", response)
}
