package main

import (
	"fmt"
	"net"

	//"github.com/vishvananda/netlink"
)

var _ fmt.Stringer = &net.IPNet{}

func main() {
	//netlink.LinkAdd(nil)
	ip, cidr, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		panic(err)
	}
	fmt.Println(ip.String())
	fmt.Printf("%+v\n", cidr)
}