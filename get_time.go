package main

import (
	"fmt"
	"log"
	"net"
	// "http"
)

func main() {
	conn, err := net.Dial("tcp", "129.6.15.28:13")
	//listener, err := net.Listen("tcp", "129.6.15.28:13")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//conn, err := listener.Accept()
		//if err != nil {
		//    fmt.Println("conn error: ", err)
		//    continue
		//}
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("conn error: ", err)
			break
		}
		fmt.Println(string(b[:n]))
	}
}
