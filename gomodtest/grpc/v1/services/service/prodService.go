package service

import (
	"context"
	"fmt"
	pb2 "gRPC/v1/services/pb"
)

type ProdService struct {

}


func (s *ProdService) GetProdStock(c context.Context, request *pb2.ProdRequest) (*pb2.ProdResponse, error)  {
	fmt.Println("传入的商品 id: ", request.ProdId)
	fmt.Println("传入的商品 购买区域: ", request.Area)

	var stock int32
	switch request.Area {
	case pb2.ProdAreas_CHINESE:
		stock = 500
	case pb2.ProdAreas_RUSSIA:
		stock = 450
	case pb2.ProdAreas_EUROPE:
		stock = 1200
	case pb2.ProdAreas_AMERICA:
		stock = 1100
	}

	return &pb2.ProdResponse{
		ProdStock: request.ProdId * stock,
	}, nil
}

func (s *ProdService) GetProdStocks(c context.Context, request *pb2.ProdRequest) (*pb2.ProdResponseList, error) {
	fmt.Println("传入的商品 id: ", request.ProdId)
	list := []*pb2.ProdResponse{
		{
			ProdStock: 123,
		},
		{
			ProdStock: 456,
		},
		{
			ProdStock: 789,
		},
	}
	p := new(pb2.ProdResponseList)
	p.List = list
	return p, nil
}

func (s *ProdService) GetProdInfo(c context.Context, request *pb2.ProdRequest) (*pb2.ProdModel, error) {
	area := request.Area
	id := request.ProdId
	fmt.Printf("获取商品信息，地区：%s, id：%d", area, id)

	var price float32
	switch request.Area {
	case pb2.ProdAreas_CHINESE:
		price = 500
	case pb2.ProdAreas_RUSSIA:
		price = 450
	case pb2.ProdAreas_EUROPE:
		price = 1200
	case pb2.ProdAreas_AMERICA:
		price = 1100
	}

	return &pb2.ProdModel{
		Id:        id,
		ProdName:  "iphone:" + area.String(),
		ProdPrice: price * 500,
	}, nil
}