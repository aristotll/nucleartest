package main

import "net"

func server() {
	net.Listen("tcp", "8888")
}
