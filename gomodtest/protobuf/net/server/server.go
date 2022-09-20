package main

import (
	"fmt"
	"net"
	"protobuftest/pb/pbfile"

	"google.golang.org/protobuf/proto"
)

func chaiBao() {
	lis, err := net.Listen("tcp", ":7788")
	assert(err, "listen error: %v")

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for {
			// b := make([]byte, 1024)	// proto: cannot parse invalid wire-format data
			b := make([]byte, 279896)
			n, err := conn.Read(b)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("read %v bytes\n", n)

			var m pbfile.Message
			if err := proto.Unmarshal(b, &m); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(m.Uid)
		}
	}
}

func zhanBao() {
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
			n, err := conn.Read(b)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("read %v bytes\n", n)

			var m pbfile.Message
			if err := proto.Unmarshal(b, &m); err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(m.Uid)
		}
	}
}

func assert(err error, format string) {
	if err != nil {
		panic(fmt.Sprintf(format, err.Error()))
	}
}

func main() {
	zhanBao()
}
