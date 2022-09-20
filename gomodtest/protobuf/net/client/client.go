package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"protobuftest/pb/pbfile"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// tcp 拆包问题
func chaiBao() {
	home := os.Getenv("HOME")
	fp := filepath.Join(home, "Downloads/jDKUARa.jpg")

	f, err := os.Open(fp)
	assert(err, "open file error: %v")
	defer f.Close()

	file, err := os.ReadFile(fp)
	assert(err, "read file error: %v")

	msg := pbfile.Message{
		Uid:         wrapperspb.UInt64(10086),
		Data:        file,
		MessageType: wrapperspb.String("image"),
	}

	conn, err := net.Dial("tcp", ":7788")
	assert(err, "dial error: %v")
	defer conn.Close()

	b, err := proto.Marshal(&msg)
	assert(err, "proto marshal error: %v")
	fmt.Printf("proto marshal size: %v\n", len(b))

	n, err := conn.Write(b)
	assert(err, "write to conn error: %v")
	fmt.Printf("write %v bytes\n", n)
}

// tcp 粘包问题
func zhanBao() {
	conn, err := net.Dial("tcp", ":7788")
	assert(err, "dial error: %v")
	defer conn.Close()

	var totalN int 
	for i := 0; i < 10; i++ {
		msg := pbfile.Message{
			Uid:         wrapperspb.UInt64(uint64(i)),
			Data:        []byte("hello"),
			MessageType: wrapperspb.String("text"),
		}

		b, err := proto.Marshal(&msg)
		assert(err, "proto marshal error: %v")
		fmt.Printf("proto marshal size: %v\n", len(b))

		n, err := conn.Write(b)
		totalN += n
		assert(err, "write to conn error: %v")
		fmt.Printf("write %v bytes\n", n)
	}
	fmt.Printf("total write %v bytes\n", totalN)
}

func assert(err error, format string) {
	if err != nil {
		panic(fmt.Sprintf(format, err.Error()))
	}
}

func main() {
	zhanBao()
}
