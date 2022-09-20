package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os/exec"
	"runtime"
)

// TODO build a rpc server in client side

type Server struct{}

func (s *Server) ExecCommand(req *Request, resp *Response) error {
	cmd := exec.Command(req.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		resp.Result = string(output)
		return fmt.Errorf("exec command [%v] error: %v", req.Command, string(output))
	}
	resp.Result = string(output)
	return nil
}

func (s *Server) GetOS(req *Request, resp *Response) error {
	resp.Result = runtime.GOOS
	return nil
}

func (s *Server) Heartbeat(req *Request, resp *Response) error {
	resp.Result = "OK"
	return nil
}

type Request struct {
	Command string
}

type Response struct {
	Result string
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
