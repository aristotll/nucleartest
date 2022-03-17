package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/ants/v2"
)

func main() {
	pool, err := ants.NewPool(1000)
	if err != nil {
		log.Fatalln(err)
	}

	count := 10000
	var wg sync.WaitGroup

	wg.Add(count)

	for i := 0; i < count; i++ {
		pool.Submit(func() {
			i := i
			defer wg.Done()
			fmt.Println(i, "run")
		})
		fmt.Printf("pool.Running(): %v\n", pool.Running())
	}

	wg.Wait()
}
