package main

import (
	"log"
	"net"
)

func main() {
	addr := &net.UDPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 8080,
	}
	// panic: listen udp :8080: address :8080: unexpected address type
	//l, err := net.Listen("udp", ":8080")
	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	for {
		b := make([]byte, 4096)
		n, addr, err := l.ReadFromUDP(b)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("read from addr: %+v, n: %v \n", addr, n)

		_, err = l.WriteToUDP(b, addr)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(string(b))
	}
}
