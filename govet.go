package main

import (
	"fmt"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9527")
	assert(err, "listen tcp error: %v", err.Error())
    _ = lis
}

func assert(err error, format string, msg string) {
	if err != nil {
		panic(fmt.Sprintf(format, msg))
	}
}
