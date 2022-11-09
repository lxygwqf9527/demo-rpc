package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"gitub.com/lxygwqf9527/demo-rpc/json_tcp/service"
)

// 约束服务端接口的实现
var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// 客户端采用json格式来编解码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	return &HelloServiceClient{
		client: client,
	}, nil
}

type HelloServiceClient struct {
	client *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, reponse *string) error {
	return c.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, reponse)
}

func main() {
	// 创建客户端
	c, err := NewHelloServiceClient("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var resp string
	if err := c.Hello("alice", &resp); err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
