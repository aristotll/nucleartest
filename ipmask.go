package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

var cidrAddr = flag.String("addr", "10.0.0.0/24", "cidr addr")

func main() {
	flag.Parse()
	// if *cidrAddr == "" {
	// 	fmt.Println("Usage: -addr 10.0.0.0/24")
	// 	return
	// }

	_, netip, err := net.ParseCIDR(*cidrAddr)
	if err != nil {
		panic(err)
	}
	netAddr := netip.IP.Mask(netip.Mask)
	fmt.Println(netAddr)

	idx := strings.Index(*cidrAddr, "/")
	b := []byte(*cidrAddr)
	copy(b[idx+1:], []byte("16"))
	*cidrAddr = string(b)

	fmt.Println(string(b), *cidrAddr)

	mask := net.CIDRMask(16, 32)
	fmt.Println(netip.IP.Mask(mask))
	fmt.Println(netip.IP.Mask(mask).Equal(netAddr))
}
