package main

import (
	"net/rpc"
	"log"
	"time"
)

type Server {}
type Request {}
type Response {}

func (s *Server) Do() {
	log.Println("start sleep")
	time.Sleep(time.Second * 10)
	log.Println("sleep done")
}

func main() {
	if err := rpc.Register(new(Server)); err != nil {
		log.Fatalln(err)
	}

	l, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalln(err)
	}

	rpc.Accept(l)
}
