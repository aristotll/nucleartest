package main

import (
	"log"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			cmd := exec.Command("./sleep")
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalln(err)
			}
		}()
	}
	wg.Wait()
}
