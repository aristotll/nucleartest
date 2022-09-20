package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	assert(err, "listen error", true)

	for {
		c, err := l.Accept()
		assert(err, "accept error", false)
		if err != nil {
			continue
		}
		go func() {
			for i := 0; i < 100; i++ {
				if i%2 == 0 {
					c.Write([]byte("aaa\n"))
				} else {
					c.Write([]byte("bbb\n"))
				}
			}
		}()
	}
}

func assert(err error, msg string, isPanic bool) {
	if err != nil {
		errmsg := fmt.Sprintf("%v: %v", msg, err)
		if isPanic {
			panic(errmsg)
		}
		log.Println(errmsg)
	}
}
