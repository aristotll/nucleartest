package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7788")
	assert(err, "listen error: %v")

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for {
			b := make([]byte, 1024)
			//time.Sleep(time.Second * 5)
			_, err := conn.Read(b)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(b))
		}
	}
}

func assert(err error, format string) {
	if err != nil {
		panic(fmt.Sprintf(format, err.Error()))
	}
}
