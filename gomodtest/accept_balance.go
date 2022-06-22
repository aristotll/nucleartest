package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"sync"
)

// 多个 goroutine accept，是负载均衡的吗
func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("start server...")
	n := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			conn, err := l.Accept()
			if err != nil {
				panic(err)
			}
			defer conn.Close()
			conn.Write([]byte(fmt.Sprintf("accept %v\n", i)))
		}(i)
	}
	wg.Wait()

// test:
// $ while true;do nc localhost 8080;done;
// accept 7
// accept 1
// accept 0
// accept 4
// accept 5
// accept 6
// accept 3
// accept 2
}
