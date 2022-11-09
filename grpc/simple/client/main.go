package main

import (
	"context"
	"fmt"

	"gitub.com/lxygwqf9527/rpc-demo/grpc/simple/server/pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.DialContext(context.Background(), "localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &pb.Request{Value: "alice"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}