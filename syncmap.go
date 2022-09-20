package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var id int64
var wg sync.WaitGroup
var m sync.Map

type A struct{ id int64 }
type B struct{ id int64 }

func NewA() *A {
	return &A{id: atomic.AddInt64(&id, 1)}
}

func NewB() *B {
	return &B{id: atomic.AddInt64(&id, 1)}
}

func main() {
	count := 10
	wg.Add(count * 2)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(NewA().id)
			if _, loaded := m.LoadOrStore(NewA().id, nil); loaded {
				panic("id 重复！")
			}
		}()

		go func() {
			defer wg.Done()
			fmt.Println(NewB().id)
			if _, loaded := m.LoadOrStore(NewA().id, nil); loaded {
				panic("id 重复！")
			}
		}()
	}
	wg.Wait()

	m.Range(func(k interface{}, v interface{}) bool {
		fmt.Printf("k: %v ", k)
		return true
	})
}
