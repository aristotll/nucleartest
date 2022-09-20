package clientFunc

import (
	"context"
	"fmt"
	pb2 "gRPC/v1/services/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"log"
)

func SubmitOrder(c *gin.Context) {
	_ = c.PostForm("id")

	client := runSpikeClient()
	stream, err := client.SubmitOrder(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	orders := make([]*pb2.SpikeUserOrder, 0)
	req := new(pb2.SpikeOrderRequest)

	go func() {
		var a int64 = 0
		for i := a; i < 10; i++ {
			orders = append(orders, &pb2.SpikeUserOrder{
				Id:          i * 234,
				No:          i * 666,
				UserId:      i,
				ProductName: "ihuawei" + string(i),
				Price:       float32(i) * 5666,
				CreateTime:  nil,
			})
			req.SpikeOrder = orders
			orders = orders[0 : 0]

			err = stream.Send(req)
			if err != nil {
				log.Fatal("send error: ", err)
			}
		}
	}()

	go func() {
		for {
			rsp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(rsp)
		}
	}()
}

func runSpikeClient() pb2.SpikeOrderServiceClient {
	// 连接到服务端
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb2.NewSpikeOrderServiceClient(conn)
	return client
}