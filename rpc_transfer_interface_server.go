package main

import (
	"net/rpc"
	"net"
	"encoding/gob"
)

// std rpc 传输接口类型，能否成功

type Interface interface {
	XX()
} 

type Service struct {}

type (
	Args struct {X Interface}
	Reply struct {Y Interface}
)

func (s *Service) Do(args *Args, reply *Reply) error {
	reply.Y = &B{"123"}
	return nil
}

type A struct {Name string}

// 实现Interface 接口
func (a *A) XX() {}

type B struct {Age string}

func (b *B) XX() {}

func main() {
	gob.Register(&A{})
	gob.Register(&B{})

	rpc.Register(new(Service))
	
	l, err := net.Listen("tcp", ":7070")
	if err != nil {
		panic(err)
	}
	
	rpc.Accept(l)
}
