package clientFunc

import (
	"context"
	"fmt"
	pb2 "gRPC/v1/services/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func CreateOrder() {
	client := runOrderClient()
	now := timestamppb.Now()
	rsp, err := client.NewOrder(context.Background(), &pb2.OrderInfo{
		Id:     123,
		No:     101359,
		UserId: 1001,
		Pay:    8391841,
		CreateTime: now,
		Detail: []*pb2.OrderDetail{
			{
				Id: 123,
				OrderId: "666",
				ProdId: 566,
				Quantity: 98765,
				TotalAmount: 999999,
			},
			{
				Id: 4365,
				OrderId: "23446",
				ProdId: 43654,
				Quantity: 23,
				TotalAmount: 56765,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp)
}

func runOrderClient() pb2.OrderServiceClient {
	// 连接到服务端
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb2.NewOrderServiceClient(conn)
	return client
}
