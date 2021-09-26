package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	for {
		buf := bufio.NewReader(conn)
		b := make([]byte, 1024)
		n, err := buf.Read(b)
		if err != nil {
			log.Println(err)
			return
		}

		wb := bufio.NewWriter(conn)
		wb.Write(b[:n])
		wb.Flush()
	}
}
