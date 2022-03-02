package main

import (
	. "syscall"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	fd, err := Socket(AF_INET, SOCK_STREAM, 0)
	if err != nil {
		t.Fatal(err)
	}
	if err := SetsockoptInt(fd, SOL_SOCKET, SO_REUSEADDR, 1); err != nil {
		t.Fatal(err)
	}
	if err := Bind(fd, &SockaddrInet4{Port: 9999, Addr: [4]byte{127, 0, 0, 1}}); err != nil {
		t.Fatal(err)
	}
	if err := Listen(fd, 1024); err != nil {
		t.Fatal(err)
	}
	for {
		connfd, _, err := Accept(fd)
		if err != nil {
			t.Log(err)
			continue
		}
		buf := make([]byte, 1024)
		_, err = Read(connfd, buf)
		if err != nil {
			t.Log(err)
			break
		}
		if _, err := Write(connfd, buf); err != nil {
			t.Log(err)
			break
		}
		Close(connfd)
	}
}

func TestClient(t *testing.T) {
	fd, err := Socket(AF_INET, SOCK_STREAM, 0)
	if err != nil {
		t.Fatal(err)
	}
	if err := Connect(fd, &SockaddrInet4{Port: 9999, Addr: [4]byte{127, 0, 0, 1}}); err != nil {
		t.Fatal(err)
	}
	if _, err := Write(fd, []byte("123")); err != nil {
		Close(fd)
		t.Fatal(err)
	}
	buf := make([]byte, 1024)
	if _, err := Read(fd, buf); err != nil {
		Close(fd)
		t.Fatal(err)
	}
	t.Log(string(buf))
	time.Sleep(time.Second * 5)
	Close(fd)
}
