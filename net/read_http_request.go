package main

import (
	"net"
	"fmt"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("new conn: ", conn.RemoteAddr().String())

		b := make([]byte, 1024)
		conn.Read(b)

		fmt.Println(string(b))
	}
}
