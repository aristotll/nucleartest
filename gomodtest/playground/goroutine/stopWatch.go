package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func watch() {
	times := time.Now()
	s := fmt.Sprintf("当前系统时间: %d:%d:%d:%d",
		times.Hour(), times.Minute(),
		times.Second(), times.Nanosecond())
	fmt.Println(s)
	defer wg1.Done()
}

func main() {
	for  {
		wg1.Add(1)
		go watch()
		wg1.Wait()
	}
}
