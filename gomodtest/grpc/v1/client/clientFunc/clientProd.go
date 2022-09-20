package clientFunc

import (
	"context"
	"fmt"
	"gRPC/v1/services/pb"
	"google.golang.org/grpc"
	"log"
)

func GetStock() {
	client := runClient()
	response, err := client.
		GetProdStock(context.Background(),
			&pb.ProdRequest{ProdId: 16, Area: pb.ProdAreas_EUROPE})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("product stocks: ", response.ProdStock)
}

// 返回一个 slice
func GetStocks()  {
	client := runClient()
	rsp, err := client.
		GetProdStocks(context.Background(),
			&pb.ProdRequest{ProdId: 32, Area: pb.ProdAreas_CHINESE})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("product stocks[] : ", rsp.List)
}

func GetProdInfo() {
	client := runClient()
	rsp, err := client.GetProdInfo(context.Background(), &pb.ProdRequest{
		ProdId: 66,
		Area:   pb.ProdAreas_CHINESE,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("product info: ", rsp)
}

func runClient() pb.ProdServiceClient {
	// 连接到服务端
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProdServiceClient(conn)
	return client
}

