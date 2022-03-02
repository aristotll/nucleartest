package main

import (
	"log"
	"net/rpc"
)

type Req struct {
	Num1 int64
	Num2 int64
}

type Rsp struct {
	Res int64
}

func main() {
	c, err := rpc.Dial("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}

	var rsp Rsp
	if err := c.Call("Server.Add", &Req{1, 2}, &rsp); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v \n", rsp)

	var b int64
	if err := c.Call("Server.XX", 10, &b); err != nil {
		log.Fatalln(err)
	}
	log.Println(b)
}
