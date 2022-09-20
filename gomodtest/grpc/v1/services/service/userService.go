package service

import (
	"context"
	"fmt"
	pb2 "gRPC/v1/services/pb"
	"io"
	"log"
)

type UserService struct {}

func (u *UserService) GetUserScore(c context.Context, us *pb2.UserScoreRequest) (*pb2.UserScoreResponse, error) {
	scores := make([]*pb2.UserScore, 0)
	for k, v := range us.UserScore {
		scores = append(scores, &pb2.UserScore{
			Id:    int32(k) + 1000,
			Score: 22 * float32(v.Id),
		})
	}
	rsp := new(pb2.UserScoreResponse)
	rsp.UserScore = scores
	return rsp, nil
}

func (u *UserService) GetUserScoreByServerStream(req *pb2.UserScoreRequest,
	stream pb2.UserScoreService_GetUserScoreByServerStreamServer) error {
	scores := make([]*pb2.UserScore, 0)
	for k, v := range req.UserScore {
		scores = append(scores, &pb2.UserScore{
			Id:    int32(k) + 1000,
			Score: 22 * float32(v.Id),
		})
		rsp := new(pb2.UserScoreResponse)
		rsp.UserScore = scores

		if (k + 1) % 2 == 0 && k > 0 {
			err := stream.Send(rsp)
			if err != nil {
				log.Fatal(err)
			}
			scores = scores[0 : 0]
		}
	}

	if len(scores) > 0 {
		rsp := new(pb2.UserScoreResponse)
		rsp.UserScore = scores
		err := stream.Send(rsp)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func (u *UserService) GetUserScoreByClientStream(stream pb2.UserScoreService_GetUserScoreByClientStreamServer) error {
	scores := make([]*pb2.UserScore, 0)
	rsp := new(pb2.UserScoreResponse)
	// 持续接收客户端传来的数据
	for {
		req, err := stream.Recv()
		// 坑: 先判断 err 是否为 eof，必须放到 err != nil 之前
		if err == io.EOF {
			return stream.SendAndClose(rsp)
		}

		if err != nil {
			return err
		}

		for k, v := range req.UserScore {
			scores = append(scores, &pb2.UserScore{
				Id:    int32(k) + 1000,
				Score: 22 * float32(v.Id),
			})
		}
		rsp.UserScore = scores
		// scores = scores [0 : 0]
	}
}

func (u *UserService) GetUserScoreByTwoWayStream(stream pb2.UserScoreService_GetUserScoreByTwoWayStreamServer) error{
	scores := make([]*pb2.UserScore, 0)
	rsp := new(pb2.UserScoreResponse)
	for {
		req, err := stream.Recv()
		fmt.Println(req)
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		for k, v := range req.UserScore {
			scores = append(scores, &pb2.UserScore{
				Id:    int32(k) + 1000,
				Score: 22 * float32(v.Id),
			})
		}
		rsp.UserScore = scores
		scores = scores [0 : 0]
		// 接收
		err = stream.Send(rsp)
		if err != nil {
			return err
		}
	}
}