package main

import (
	"fmt"
	"sync"
)

type queue []int

type syncq struct {
	mu   sync.Mutex
	q    queue
	size int64
	wg   sync.WaitGroup
}

func NewSyncq(size int64) *syncq {
	return &syncq{
		size: size,
	}
}

func (s *syncq) remove() {
	defer s.wg.Done()
	s.mu.Lock()
	for len(s.q) > 0 {
		s.q = s.q[1:]
		fmt.Println("remove success: ", s.q)
		s.mu.Unlock()
		return
	}
	fmt.Println("remove fail")
	s.mu.Unlock()
}

func (s *syncq) add(v int) {
	defer s.wg.Done()
	s.mu.Lock()
	for (int64)(len(s.q)) < s.size {
		s.q = append(s.q, v)
		fmt.Println("add success: ", s.q)
		s.mu.Unlock()
		return
	}
	fmt.Println("add fail")
	s.mu.Unlock()
}

func main() {
	q := NewSyncq(10)
	num := 100

	q.wg.Add(num)
	for i := 0; i < num; i++ {
		go q.add(i)
		go q.remove()
	}
	q.wg.Wait()
}
