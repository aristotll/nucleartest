package main

import (
	"log"
	"net"
)

func handler(conn net.Conn) error {
	b := make([]byte, 1024)
	_, err := conn.Read(b)
	if err != nil {
		return err
	}

	_, err = conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			break
		}

		go func() {
			if err := handler(conn); err != nil {
				log.Println(err)
			}
		}()
	}
}
