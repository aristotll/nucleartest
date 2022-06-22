package main

import (
	"log"
	. "syscall"
)

const eventNum = 10

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
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

	kfd, err := Kqueue()
	if err != nil {
		panic(err)
	}

	change := Kevent_t{
		Ident:  uint64(sfd),
		Filter: EVFILT_READ,
		Flags:  EV_ADD | EV_ENABLE,
		Fflags: 0,
		Data:   0,
		Udata:  nil,
	}
	_, err = Kevent(kfd, []Kevent_t{change}, nil, nil) // 添加监听事件，第三个参数传空
	if err != nil {
		panic(err)
	}

	for {
		events := make([]Kevent_t, eventNum)
		// 获取就绪事件，第二个参数传空
		readyEventNum, err := Kevent(kfd, nil, events, nil) // 无限等待直到有事件产生
		//fmt.Printf("readyEventNum: %v\n", readyEventNum)
		if err != nil && err != EINTR {
			panic(err)
		}
		for i := 0; i < readyEventNum; i++ {
			event := events[i]
			log.Printf("event fd: %v \n", event.Ident)
			efd := int(event.Ident)
			if efd == sfd {
				nfd, _, err := Accept(sfd)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Printf("accept event, fd: %v\n", nfd)
				if err := SetNonblock(nfd, true); err != nil {
					log.Println(err)
					continue
					//panic(err)
				}
				change := Kevent_t{
					Ident:  uint64(nfd),
					Filter: EVFILT_READ,
					Flags:  EV_ADD | EV_ENABLE,
					Fflags: 0,
					Data:   0,
					Udata:  nil,
				}
				// 添加监听事件，第三个参数传空
				_, err = Kevent(kfd, []Kevent_t{change}, nil, nil)
				if err != nil {
					panic(err)
				}
			} else {
				log.Printf("read event\n")
				b := make([]byte, 1024)
				_, err := Read(efd, b)
				if err != nil {
					log.Println(err)
					Close(efd)
					continue
				}
				_, err = Write(efd, b)
				if err != nil {
					log.Println(err)
					Close(efd)
					continue
				}
			}
		}
	}
}
