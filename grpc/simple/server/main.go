package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"gitub.com/lxygwqf9527/rpc-demo/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

// type HelloServiceServer interface {
// 	Hello(context.Context, *Request) (*Response, error)
// 	mustEmbedUnimplementedHelloServiceServer()
// }

type HelloServiceServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServiceServer) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Value: fmt.Sprintf("hello %s", req.Value)}, nil
}

func main() {
	// s grpc.ServiceRegistrar grpc Server
	// srv HelloServiceServer
	server := grpc.NewServer()
	// 把实现类注册给GRPC Server
	pb.RegisterHelloServiceServer(server, new(HelloServiceServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	log.Printf("grpc listen addr: 127.0.0.1:1234")
	// 监听Socket,HTTP2内置
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
