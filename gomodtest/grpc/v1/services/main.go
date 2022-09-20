package main

import (
	pb2 "gRPC/v1/services/pb"
	service2 "gRPC/v1/services/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	pb2.RegisterProdServiceServer(server, new(service2.ProdService))
	pb2.RegisterOrderServiceServer(server, new(service2.OrderService))
	pb2.RegisterUserScoreServiceServer(server, new(service2.UserService))
	pb2.RegisterSpikeOrderServiceServer(server, new(service2.SpikeService))

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}
