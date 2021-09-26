package main

import (
	"encoding/gob"
	"log"
	"net"
)

type awk struct {
	Str string
	Par string
}

func main() {
	o := &awk{
		Str: "abcd",
		Par: "a",
	}

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := listen.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	if err := gob.NewEncoder(conn).Encode(o); err != nil {
		log.Println(err)
		return
	}

	var o1 awk
	if err := gob.NewDecoder(conn).Decode(&o1); err != nil {
		log.Println(err)
		return
	}
	log.Println(o1)
}
