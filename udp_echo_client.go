package main

import (
	"log"
	"net"
)

func main() {
	raddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// writeToUDP：use of WriteTo with pre-connected connection
	b := make([]byte, 2000)
	for i := 0; i < len(b); i++ {
		b[i] = 'a'
	}
	_, err = conn.Write(b)
	if err != nil {
		panic(err)
	}

	b = make([]byte, 4096)
	_, err = conn.Read(b)
	if err != nil {
		panic(err)
	}
	log.Println(len(b), string(b))
}
