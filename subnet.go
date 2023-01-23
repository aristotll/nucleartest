package main

import (
	"fmt"
	"net"
)

func main() {
    _, ipnet, err := net.ParseCIDR("192.168.110.0/24")
    if err != nil {
        panic(err)
    }
    fmt.Println(ipnet.IP, ipnet.Mask)
    i, b := ipnet.Mask.Size()
    fmt.Println(i, b)

    // Output:
    // 192.168.110.0 ffffff00
    // 24 32
}
