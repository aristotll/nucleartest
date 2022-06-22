package main

import (
	"net"
	"log"
)

func main() {
	addr := ":8080"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("listen error: ", err)	
	}
	log.Printf("listen in %v \n", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			break
		}
		b := make([]byte, 1024)
		for {
			_, err := conn.Read(b)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println(string(b))
		}
	}
}
