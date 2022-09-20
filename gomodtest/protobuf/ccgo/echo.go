//go:build ignore

package main

import (
	"github.com/golang/protobuf/proto"
	pb "gomodtest/protobuf/proto"
	"log"
)

func main() {
	sct := &pb.Request{Message: "echo testgeneric"}
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
