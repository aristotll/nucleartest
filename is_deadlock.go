package main

import (
	"log"
	"time"
)

func main() {
	var ch chan struct{}
	ch1 := make(chan int)
	ch1 <- 1

	go func() {
		for v := range ch1 {
			log.Println(v)
			time.Sleep(time.Second)
		}
	}()

	log.Println("wait")
	<-ch
}
