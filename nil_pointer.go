package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7788")
	assert(err, "listen error: %v", err.Error())
    _ = lis
}

func assert(err error, format, msg string) {
	if err != nil {
		panic(fmt.Sprintf(format, msg))
	}
}
