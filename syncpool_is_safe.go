package main

import (
	"sync"
	"fmt"
)

type Struct struct {
	name string
}

func main() {
	var wg sync.WaitGroup
	p := &sync.Pool{
		New: func() interface{} {
				return &Struct{name: "zhang3"}
		},
	}

	var count = 100

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			v := p.Get()
			fmt.Printf("goroutine[%d] get: %v \n", i, v)
			p.Put(v)
		}()
	}

	wg.Wait()
}
