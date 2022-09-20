package clientFunc

import (
	"context"
	"fmt"
	pb2 "gRPC/v1/services/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

// 批量传入用户 id，并获取对应积分
func GetUserScores() {
	client := runUserClient()
	scores := make([]*pb2.UserScore, 5)
	for i := 0; i < 5; i++ {
		scores[i] = &pb2.UserScore{
			Id:    int32(i) + 1000,
			Score: 0,
		}
	}

	req := new(pb2.UserScoreRequest)
	req.UserScore = scores
	score, err := client.GetUserScore(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user scores: ", score)
}

func GetUserScoresByServerStream() {
	client := runUserClient()
	scores := make([]*pb2.UserScore, 5)
	for i := 0; i < 5; i++ {
		scores[i] = &pb2.UserScore{
			Id:    int32(i) + 1000,
			Score: 0,
		}
	}

	req := new(pb2.UserScoreRequest)
	req.UserScore = scores

	stream, err := client.GetUserScoreByServerStream(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("服务器流处理的数据：")
	for {
		recv, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(recv.UserScore)
	}

}

func GetUserScoresByClientStream() {
	client := runUserClient()
	stream, err := client.GetUserScoreByClientStream(context.Background())
	if err != nil {
		log.Fatal("get stream error: ", err)
	}

	scores := make([]*pb2.UserScore, 0)
	req := new(pb2.UserScoreRequest)
	for i := 0; i < 5; i++ {
		scores = append(scores, &pb2.UserScore{
			Id:    5 * int32(i),
			Score: 8560 * float32(i),
		})
		req.UserScore = scores
		err := stream.Send(req)
		scores = scores[0 : 0]
		if err != nil {
			log.Fatal("send err: ", err)
		}
	}
	// 说明数据已经全部发送完成了，可以接收服务端的响应数据了
	recv, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("recv error: ", err)
	}
	fmt.Println(recv.UserScore)
}

// 双向流
func GetUserScoresByTwoWayStream() {
	client := runUserClient()
	stream, err := client.GetUserScoreByTwoWayStream(context.Background())
	if err != nil {
		log.Fatal("get stream error: ", err)
	}

	scores := make([]*pb2.UserScore, 0)
	req := new(pb2.UserScoreRequest)
	for i := 0; i < 5; i++ {
		scores = append(scores, &pb2.UserScore{
			Id:    5 * int32(i),
			Score: 8560 * float32(i),
		})
		req.UserScore = scores

		err := stream.Send(req)
		scores = scores[0 : 0]
		if err != nil {
			log.Fatal("send err: ", err)
		}

		rsp, err := stream.Recv()
		if err == io.EOF {
			// 接收完成，跳出死循环
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(rsp)
	}
}


func runUserClient() pb2.UserScoreServiceClient {
	// 连接到服务端
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb2.NewUserScoreServiceClient(conn)
	return client
}


