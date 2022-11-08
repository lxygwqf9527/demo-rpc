package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"gitub.com/lxygwqf9527/rpc-demo/rpc_interface/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloService)(nil)

type HelloService struct {
}

func (s *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello, %s", request)
	return nil

}

func main() {
	// 把rpc对外暴露的对象注册到rpc框架内部
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})
	// 1.先监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 每个客户端单独启用一个routine来处理
		go rpc.ServeConn(conn)
	}
}
