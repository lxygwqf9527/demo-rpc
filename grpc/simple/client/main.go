package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lxygwqf9527/demo-rpc/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.DialContext(context.Background(), "localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)

	// req <--> resp
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "alice"})
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
