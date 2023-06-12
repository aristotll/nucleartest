package main

import (
	"fmt"
	"net"
)

func main() {
	ip, ipv4Net, err := net.ParseCIDR("192.168.31.130/25")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取 CIDR IP 的网络地址
	fmt.Println(ip, ipv4Net.String())
}
