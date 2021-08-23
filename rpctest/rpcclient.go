package main

import (
	."firstGo/rpctest/utils"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var serverAddress = "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":8080")	//与服务器建立连接
	if err != nil {
		log.Fatalf("建立与服务端连接失败:", err)
	}

	//client1(client)
	clien2(client)


}

/**
* client1
* @Description: 同步调用
* @param cli
*/
func client1(cli *rpc.Client)  {
	args := &Args{A:8, B:4}
	var reply int
	err := cli.Call("MathServer.Multiply", args, &reply)	//调用服务端注册的rpc方法
	if err != nil {
		log.Fatal("调用远程方法 MathService.Multiply 失败:", err)
	}
	fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)
}

/**
* clien2
* @Description: 异步调用
* @param cli
*/
func clien2(cli *rpc.Client)  {
	args := &Args{A:8, B:4}
	var reply int

	dicideCall := cli.Go("MathServer.Divide", args, &reply, nil)
	for  {
		select {
		case <-dicideCall.Done:
			fmt.Printf("%d/%d=%d\n", args.A, args.B, reply)
			return
		}
	}
}