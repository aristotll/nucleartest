package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn,
			time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 如果不添加 go ，则只有一条线程执行，一次只能处理一个客户需求
		// 直到第一个客户端执行结束，第二个客户端才能执行
		go handler(conn)
	}

}
