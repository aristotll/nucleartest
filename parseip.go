package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("localhost:8080")
	fmt.Println(ip)
}
