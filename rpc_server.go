package main

import (
	"log"
	"net"
	"net/rpc"
)

type Server struct{}

type Req struct {
	Num1 int64
	Num2 int64
}

type Rsp struct {
	Res int64
}

func (s *Server) Add(r *Req, rr *Rsp) error {
	rr.Res = r.Num1 + r.Num2
	return nil
}

func (s *Server) Sub(r *Req, rr *Rsp) error {
	rr.Res = r.Num1 - r.Num2
	return nil
}

func (s *Server) Mul(r *Req, rr *Rsp) error {
	rr.Res = r.Num1 * r.Num2
	return nil
}

func (s *Server) Div(r *Req, rr *Rsp) error {
	rr.Res = r.Num1 / r.Num2
	return nil
}

func main() {
	if err := rpc.Register(&Server{}); err != nil {
		log.Fatalln(err)
	}

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go rpc.ServeConn(conn)
	}

}
