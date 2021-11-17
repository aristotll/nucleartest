package main

import (
	//"runtime"
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func () {
		defer wg.Done()
		var i int64
		for {
			i++
			if i == 100 {
				//runtime.Goexit()
				return
			}
			fmt.Printf("i: %d \n", i)
		}
		
	}()

	wg.Wait()
}