package main

import (
	"log"
	"net"
	"time"
)

func main() {
	addr := ":8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln("dial error: ", err)
	}

	b := make([]byte, 4096)
	for {
		_, err := conn.Read(b)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(b))
		time.Sleep(time.Second)
	}
}
