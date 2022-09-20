package main

import (
	"log"
	"syscall"
)

func main() {
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln(err)
	}

	if err := syscall.Connect(sockfd, &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}); err != nil {
		log.Fatalln(err)
	}

	for {

	}
}
