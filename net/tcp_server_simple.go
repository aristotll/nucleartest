package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		log.Fatalln("listen error: ", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			break
		}
		b := make([]byte, 1024)
		io.ReadFull(conn, b[:2])
		log.Println(string(b[0]), string(b[1]))

		io.ReadFull(conn, b[:10])
		log.Println(string(b[:10]))

		log.Println(string(b))
	}
}
