package main

import (
	"log"
	"net"
	"net/rpc"
	"strings"
)

type Server struct{}

func (s *Server) Call(req *Request, resp *Response) error {
	var sb strings.Builder
	sb.WriteString(req.A)
	sb.WriteString(req.B)
	sb.WriteString(req.C)
	sb.WriteString(req.D)
	sb.WriteString(req.E)
	sb.WriteString(req.F)
	sb.WriteString(req.G)
	resp.Str = sb.String()
	return nil
}

type Request struct {
	A, B, C, D, E, F, G string
}

type Response struct {
	Str string
}

func main() {
	if err := rpc.Register(new(Server)); err != nil {
		log.Println(err)
		return
	}

	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
		return
	}

	rpc.Accept(l)
}
