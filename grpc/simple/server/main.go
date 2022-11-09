package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/lxygwqf9527/demo-rpc/grpc/simple/server/pb"
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
	return &pb.Response{Value: fmt.Sprintf("hello, %s", req.Value)}, nil
}
func (s *HelloServiceServer) Channel(stream pb.HelloService_ChannelServer) error {
	for {
		// 接收请求

		req, err := stream.Recv()
		if err != nil {
			log.Printf("recv error, %s", err)
			if err == io.EOF {
				log.Printf("client closed")
				return nil
			}
			return err
		}
		resp := &pb.Response{Value: fmt.Sprintf("hello Channel, %s", req.Value)}
		// 响应请求
		err = stream.Send(resp)
		if err != nil {
			if err == io.EOF {
				log.Printf("client closed")
				return nil
			}
			return err
		}
	}
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
