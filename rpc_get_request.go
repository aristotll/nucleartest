package main

import (
	"fmt"
)

type Server struct {
	freeReq *Request
}

type Request struct {
	next *Request
}

func (s *Server) getRequest() *Request {
	req := s.freeReq
	if req == nil {
		req = new(Request)
	} else {
		s.freeReq = req.next
		*req = Request{}
	}
	return req
}

func main() {
	s := new(Server)
	count := 10
	for i := 0; i < count; i++ {
		v := s.getRequest()
		fmt.Println(v)
	}
}
