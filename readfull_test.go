package main

import (
	"io"
	"log"
	"net"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestServer929(t *testing.T) {
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
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	var buf []byte
	size := 9
	for {
		b := make([]byte, size)
		n, err := io.ReadFull(conn, b)
		if err != nil && err != io.EOF {
			log.Println(err)
			buf = append(buf, b[:n]...)
			size -= len(buf)
			continue
		}

		if len(buf) > 0 {
			buf = append(buf, b...)
			//b = append(b, buf...)
			b = buf
		}

		_, err = conn.Write(b)
		if err != nil {
			log.Println(err)
			break
		}
	}
	log.Println("done")
}

func TestClient929(t *testing.T) {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	conn.Write([]byte("123"))
	conn.Write([]byte("456789"))

	//for {
	//	b := make([]byte, 1024)
	//	_, err = conn.Read(b)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Println(string(b))
	//}

	b := make([]byte, 1024)
	_, err = conn.Read(b)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))

	time.Sleep(time.Second * 10)
}
