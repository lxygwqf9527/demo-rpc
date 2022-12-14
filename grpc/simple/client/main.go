package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lxygwqf9527/demo-rpc/grpc/middleware/client"
	"github.com/lxygwqf9527/demo-rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 添加认证信息
	crendential := client.NewAuthentication("admin", "123456")
	conn, err := grpc.DialContext(context.Background(), "localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(crendential))
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)

	// req <--> resp
	// 添加凭证信息

	// crendential := server.NewClientCredential("admin", "123456")
	// ctx := metadata.NewOutgoingContext(context.Background(), crendential)

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs())
	resp, err := client.Hello(ctx, &pb.Request{Value: "bob"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// stream
	stream, err := client.Channel(ctx)
	if err != nil {
		panic(err)
	}
	// 启动一个Goroutine来发送请求
	go func() {
		for {
			err := stream.Send(&pb.Request{Value: "alice"})
			if err != nil {
				panic(err)
			}
			time.Sleep(3 * time.Second)
		}

	}()

	for {
		// 主循环 负责接受响应
		resp, err = stream.Recv()

		if err != nil {
			panic(err)
		}

		fmt.Println(resp)
	}

}
