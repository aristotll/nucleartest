package main

import (
	"fmt"
	"log"
	. "syscall"
)

const eventNum = 10

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

// BUG 记录：
// 1. 客户端断开连接后，服务端仍然会产生读事件
// 客户端断开连接后，服务端的 Read 会发生 connection reset by peer 错误并进入错误处理分支， Close 掉这
// 个连接，但是因为错误处理采用的是 continue 而不是 break（因为可能会同时产生多个事件，也就是 kevent 的返
// 回值，如果使用 break，会导致后面的事件全部被放弃处理，因为多个事件中的某一个事件产生错误，而直接跳过后续事
// 件的处理，显然是不合理的），会重新进入外层的死循环，又因为会继续产生读事件，导致 Kevent 函数成功返回，进
// 入到读事件分支，进行 Read 操作，此时会报错 bad file descriptor，continue 到最外层死循环，如此反复，
// 直到循环很多次以后会抛出一个 panic


// BUG 记录：
// 1. 客户端断开连接后，服务端仍然会产生读事件，此时 Read 会发生 connection reset by peer 错误并 Close 掉
// 连接，但是仍然会发生读事件，导致 Kevent 成功返回，然后 setNonBlock(event.Ident) 会报错：bad file descriptor
// 如果遇到错误直接 break 不用 continue，那么其他的事件就不能处理了，因为在死循环里，所以会一直执行 setNonBlock，
// 一直报错 bad file descriptor，直到 inappropriate ioctl for device（频率过快），此时会程序会被阻塞
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

	var (
		changes = make([]Kevent_t, 0) // 监听列表
		events  = make([]Kevent_t, eventNum) // 发生的事件列表，len 不能为 0
	)

	kfd, err := Kqueue()
	if err != nil {
		panic(err)
	}

	readEvent := Kevent_t{
		Ident:  uint64(sfd),
		Filter: EVFILT_READ,
		Flags:  EV_ADD | EV_ENABLE,
		Fflags: 0,
		Data:   0,
		Udata:  nil,
	}
	changes = append(changes, readEvent)
	// BUG1 解决步骤1 ：新增下面这个函数调用
	// _, err = Kevent(kfd, changes, nil, nil)
	// if err != nil {
	// 	panic(err)
	// }

	for {
		// BUG1 解决步骤2：
		// 这里的第二个参数必须传 nil，如果传 changes 就会出现 BUG1 的情况
		readyEventNum, err := Kevent(kfd, changes, events, nil) // 无限等待直到有事件产生
		fmt.Printf("readyEventNum: %v\n", readyEventNum)
		if err != nil && err != EINTR {
			panic(err)
		}
		log.Printf("events: %v \n", events)
		log.Printf("changes: %v \n", changes)
		log.Printf("readyEventNum: %v\n", readyEventNum)
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
				changes = append(changes, Kevent_t{
					Ident:  uint64(nfd),
					Filter: EVFILT_READ,
					Flags:  EV_ADD | EV_ENABLE,
					Fflags: 0,
					Data:   0,
					Udata:  nil,
				})
				// BUG1 解决步骤3 ：新增下面这个函数调用
				// _, err = Kevent(kfd, changes, nil, nil)
				// if err != nil {
				// 	panic(err)
				// }
			} else /*if event.Filter == EVFILT_READ*/ {
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
