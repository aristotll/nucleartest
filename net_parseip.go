package main

import (
    "fmt"
    "net"
)

func main() {
    ip := net.ParseIP("127.0.0.1")
    fmt.Println(ip.String())
}
