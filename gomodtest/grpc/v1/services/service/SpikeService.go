package service

import (
	"fmt"
	"gRPC/v1/services/pb"
	"io"
	"log"
)

type SpikeService struct {

}

func (s *SpikeService) SubmitOrder(stream pb.SpikeOrderService_SubmitOrderServer) error {
	request := make(chan *pb.SpikeOrderRequest)
	go func() {
		for {
			req, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal("recv error: ", err)
			}

			request <- req

			fmt.Println("服务端接收到的数据: ", req)
		}
	}()

	go func() {
		msg := make([]string, 0)
		rsp := new(pb.SpikeOrderResponse)

		for v := range request {
			// fmt.Println("v: ", v)
			for _, val := range v.SpikeOrder {
				msg = append(msg, "订购成功，商品信息: " + val.ProductName)
			}
		}
		fmt.Println("msg: ", msg)
		// msg = append(msg, "订购成功，商品名: ", )

		rsp.Messages = msg
		err := stream.Send(rsp)
		if err != nil {
			log.Fatal("send error:  ", err)
		}
	}()
	return nil
}