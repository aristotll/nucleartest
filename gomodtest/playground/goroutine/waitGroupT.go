package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mx sync.Mutex

func handle(i int) {
	fmt.Println("func run", i)
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go handle(i)
	}
	fmt.Println("main run")
	wg.Wait()
}




