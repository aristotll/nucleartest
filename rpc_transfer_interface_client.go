package main

import (
	"net/rpc"
	"fmt"
	"encoding/gob"
)

type Interface interface {
	XX()
}

type (
	Args struct {X Interface}
	Reply struct {Y Interface}
)

type A struct {Name string}

// 实现Interface 接口
func (a *A) XX() {}

type B struct {Age string}

func (b *B) XX() {}

func main() {
	gob.Register(&A{})
	gob.Register(&B{})	

	c, err := rpc.Dial("tcp", ":7070")
	if err != nil {
		panic(err)
	}


	var (
		args = &Args{X: &A{"wang"}}
		reply = &Reply{}
	)

	if err := c.Call("Service.Do", args, reply); err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", reply)
}
