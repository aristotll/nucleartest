package service

import (
	"context"
	"fmt"
	pb2 "gRPC/v1/services/pb"
)

type OrderService struct {

}

func (o *OrderService) NewOrder(c context.Context, req *pb2.OrderInfo) (*pb2.OrderCreateResponse, error) {
	fmt.Println("提交的订单信息: ", req)
	// fmt.Println(req.Detail)
	return &pb2.OrderCreateResponse{
		Status:  "ok",
		Message: "create order successfully",
	}, nil
}
