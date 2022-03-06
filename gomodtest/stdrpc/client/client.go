package main

import (
	"log"
	"net/rpc"
)

type Request struct {
	Command string
}

type Response struct {
	Result string
}

func main() {
	client, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	req := new(Request)
	req.Command = "backup"
	resp := new(Response)

	if err := client.Call("Server.Backup", req, resp); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v \n", resp)
}
