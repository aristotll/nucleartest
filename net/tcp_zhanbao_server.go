package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 10)
		for {
			_, err = conn.Read(b)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println(string(b))
		}
	}
}
