package main

import (
	"context"
	"fmt"
	"gRPC/v2/_go/proto"
	"google.golang.org/grpc"
	"log"
)

type GreeterClient struct {

}

func (g *GreeterClient) SayHello(ctx context.Context, in *proto.HelloRequest, opts ...grpc.CallOption) (*proto.HelloReply, error) {
	panic("implement me")
}

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	reply, err := proto.
		NewGreeterClient(conn).
		SayHello(context.Background(), &proto.HelloRequest{Name: "zhang3"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply)
}
