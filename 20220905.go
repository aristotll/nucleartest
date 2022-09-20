package main

import (
    "net"
)

func main() {
    lis, err := net.Listen("tcp", ":7788")
    if err != nil {
        panic(err)
    }
    lis.Accept()
}
