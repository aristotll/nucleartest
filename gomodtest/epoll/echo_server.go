package main

import (
	"log"
	. "syscall"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func main() {
	sfd, err := Socket(AF_INET, SOCK_STREAM, 0)
	if err != nil {
		panic(err)
	}
	log.Printf("listen fd: %v\n", sfd)

	if err := Bind(sfd, &SockaddrInet4{Port: 8080, Addr: [4]byte{127, 0, 0, 1}}); err != nil {
		panic(err)
	}

	if err := Listen(sfd, 1024); err != nil {
		panic(err)
	}

	epfd, err := EpollCreate(10)
	if err != nil {
		panic(err)
	}

	if err := EpollCtl(epfd, EPOLL_CTL_ADD, sfd, &EpollEvent{
		Events: EPOLLIN,
		Fd:     int32(sfd),
		Pad:    0,
	}); err != nil {
		panic(err)
	}

	for {
		events := make([]EpollEvent, 20)
		n, err := EpollWait(epfd, events, -1)
		if err != nil {
			log.Println(err)
			continue
		}
		for i := 0; i < n; i++ {
			event := events[i]
			if event.Fd == int32(sfd) {
				nfd, _, err := Accept(sfd)
				if err != nil {
					log.Println(err)
					continue
				}
				if err := EpollCtl(epfd, EPOLL_CTL_ADD, nfd, &EpollEvent{
					Events: EPOLLIN,
					Fd:     int32(nfd),
					Pad:    0,
				}); err != nil {
					log.Println(err)
					continue
				}
			} else {
				b := make([]byte, 1024)
				_, err := Read(int(event.Fd), b)
				if err != nil {
					log.Println(err)
					Close(int(event.Fd))
					continue
				}
				_, err = Write(int(event.Fd), b)
				if err != nil {
					log.Println(err)
					Close(int(event.Fd))
					continue
				}
			}
		}
	}
}
