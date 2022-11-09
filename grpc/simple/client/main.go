package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lxygwqf9527/demo-rpc/grpc/middleware/server"
	"github.com/lxygwqf9527/demo-rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	conn, err := grpc.DialContext(context.Background(), "localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)

	// req <--> resp
	// 添加凭证信息

	crendential := server.NewClientCredential("admin", "123456")
	ctx := metadata.NewOutgoingContext(context.Background(), crendential)
	resp, err := client.Hello(ctx, &pb.Request{Value: "alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	// stream
	stream, err := client.Channel(context.Background())
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
