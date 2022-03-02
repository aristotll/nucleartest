package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":6666")
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 100; i++ {
		//si := strconv.Itoa(i)
		n, err := conn.Write([]byte("aaaaa"))
		if err != nil {
			log.Printf("write error: %v \n", err)
			return
		}
		log.Printf("write %d bytes \n", n)
		// server 缓冲区为 10，这里先发送 5，再睡眠 3s，看看对端会不会等待
		//time.Sleep(time.Second * 3)
	}
}
