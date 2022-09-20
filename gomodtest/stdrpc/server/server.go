package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type Request struct {
	Command string
}

type Response struct {
	Result string
}

type Server struct{}

func (s *Server) Backup(req *Request, resp *Response) error {
	cmd := req.Command
	log.Printf("接收到 [%v] 命令 \n", cmd)
	log.Println("执行中....")

	time.Sleep(time.Second * 5)

	resp.Result = "ok"
	log.Println("执行完成")

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	if err := rpc.Register(new(Server)); err != nil {
		log.Fatalln(err)
	}

	rpc.Accept(listen)

	//for {
	//	conn, err := listen.Accept()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	rpc.Accept()
	//	go rpc.ServeConn(conn)
	//}
}
