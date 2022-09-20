package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var cmd = flag.String("c", "", "input a command")
var port = flag.String("p", "", "input port")
var host = flag.String("h", "", "input host")

type Request struct {
	Command string
}

type Response struct {
	Result string
}

func main() {
	flag.Parse()
	if *cmd == "" {
		log.Fatalln("usage: go run xxx.go -c [command]")
	}

	addr := fmt.Sprintf("%s:%s", *host, *port)

	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	var (
		req  = &Request{Command: *cmd}
		resp = &Response{}
	)

	if err := client.Call("Server.GetOS", req, resp); err != nil {
		log.Fatalln(err)
	}
	log.Printf("this machine OS is %v\n", resp.Result)

	if err := client.Call("Server.ExecCommand", req, resp); err != nil {
		log.Fatalln(err)
	}

	log.Printf("result: \n %+v \n", string(resp.Result))
}
