package main

import (
	"net"
	"log"
	"time"
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
		log.Printf("client %v in\n", conn.RemoteAddr())
		for {
			_, err := conn.Write([]byte("123"))
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("write ok")
			time.Sleep(time.Second)
		}
	}
}
