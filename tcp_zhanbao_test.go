package main

import (
	"log"
	"net"
	"strconv"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestServer929_(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 10)
		for {
			_, err = conn.Read(b)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println(string(b))
		}
	}
}

func TestClient929_(t *testing.T) {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 100; i++ {
		si := strconv.Itoa(i)
		n, err := conn.Write([]byte(si + "is come back [done]"))
		if err != nil {
			log.Printf("write error: %v \n", err)
			return
		}
		log.Printf("write %d bytes \n", n)
	}
}
