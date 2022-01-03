package main

import (
	"github.com/golang/protobuf/proto"
	pb "gomodtest/protobuf/proto"
	"log"
	"strings"
)

func main() {
	pb.NewEchoClient()
	sct := &pb.Request{Message: "echo test"}
	b, err := proto.Marshal(sct)
	if err != nil {
		log.Fatalln(err)
	}

	o := &pb.Request{}
	if err := proto.Unmarshal(b, o); err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", o)
}


