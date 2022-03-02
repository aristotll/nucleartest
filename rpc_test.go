package main

import (
	"log"
	"net"
	"net/rpc"
	"sync"
	"testing"
)

type Struct struct{}

func (s *Struct) XX(x, y *int64) error {
	*y = *x * 10
	return nil
}

func Test1server(t *testing.T) {
	if err := rpc.Register(new(Struct)); err != nil {
		t.Fatal(err)
	}
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		t.Fatal(err)
	}
	rpc.Accept(lis)
}

func Test1client(t *testing.T) {
	count := 1000
	var wg sync.WaitGroup
	wg.Add(count)
	// 10 个 goroutine，每个 goroutine 不断循环调用 rpc
	for i := 0; i < count; i++ {
		go func(i int64) {
			defer wg.Done()
			cli, err := rpc.Dial("tcp", ":9999")
			if err != nil {
				log.Fatal(err)
			}
			for j := 0; j < 10; j++ {
				a := 1
				var b int64
				if err := cli.Call("Struct.XX", a, &b); err != nil {
					log.Fatal(err)
				}
				t.Logf("[goroutine%d]%+v", i, b)
			}
		}(int64(i))
	}
	wg.Wait()
}
