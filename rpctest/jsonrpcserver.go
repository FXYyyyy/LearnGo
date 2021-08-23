package main

import (
	. "firstGo/rpctest/utils"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type ServiceHandler struct {

}

func (sh *ServiceHandler) GetName(id int, item *Item) error {
	log.Printf("receive GetName call, id: %d", id)
	item.Id = id
	item.Name = "zhizhi"
	return nil
}

func (sh *ServiceHandler) SaveName(item Item, resp *Response) error {
	log.Printf("receive SaveName call, Item: %v", item)
	resp.Ok = true
	resp.Id = item.Id
	resp.Message = "save success"
	return nil
}

func main() {
	//初始化rpc服务端
	server := rpc.NewServer()

	//监听端口8080
	listener,err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("监听端口失败：%v", err)
	}

	defer listener.Close()	//关闭监听

	log.Println("Start listen on port localhost:8080")

	sh := new(ServiceHandler)
	err = server.Register(sh)	//注册rpc对象
	if err != nil {
		log.Fatalf("注册服务处理器失败：%v", err)
	}

	//等待并处理客户端的请求
	for  {
		conn,err := listener.Accept()
		if err != nil {
			log.Fatalf("接收客户端连接请求失败: %v", err)
		}

		//自定义 RPC 编码器：新建一个 jsonrpc 编码器，并将该编码器绑定给 RPC 服务端处理器
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
