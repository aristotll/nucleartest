package main

import (
	"log"
	"syscall"
	"time"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln(err)
	}

	if err := syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}); err != nil {
		log.Fatalln(err)
	}

	if err := syscall.Listen(fd, 1); err != nil {
		log.Fatalln(err)
	}

	log.Println("开始睡眠")
	time.Sleep(time.Second * 60)
	log.Println("睡眠结束")

	for {
		connfd, _, err := syscall.Accept(fd)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(connfd)
	}
}
