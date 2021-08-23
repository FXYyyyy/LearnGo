package main

import (
	"errors"
	. "firstGo/rpctest/utils"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MathServer struct {
	
}

func (m *MathServer) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (m *MathServer) Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}

	*reply = args.A/args.B
	return nil
}

func main() {
	math := new(MathServer)
	rpc.Register(math)	//注册math对象

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8080")	//监听
	if err != nil {
		log.Fatalf("启动服务监听失败：%v", err)
	}

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("启动 HTTP 服务失败:", err)
	}
}
