package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

var conns = make([]net.Conn, 0, 10)

func main() {
	addr := ":8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Println("listen in ", addr)
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		conns = append(conns, conn)
		for _, c := range conns {
			go ServerConn(c)
		}
	}
}

func ServerConn(conn net.Conn) {
	for {
		buf := make([]byte, 4096)
		_, err := conn.Read(buf)
		if err != nil {
			log.Printf("[%v]read error: %v\n", conn.RemoteAddr(), err)
			break
		}
		for _, c := range conns {
			if c == conn {	// 不需要给自己发
				continue
			}
			// 写入格式：[ip]消息
			b := bytes.NewBuffer(make([]byte, 0))
			b.WriteString(fmt.Sprintf("[%v]", conn.RemoteAddr()))
			b.Write(buf)
			_, err := c.Write(b.Bytes())
			if err != nil {
				log.Printf("%v write to %v error: %v\n", conn.RemoteAddr(), c.RemoteAddr(), err)
				continue
			}
		}
	}
}