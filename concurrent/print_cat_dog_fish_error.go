package main

import (
	"fmt"
	//"time"
	//"sync/waitgroup"
)

// 三条线程交替打印 dog fish cat 100 次
func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("dog")
			ch1 <- struct{}{}
		}()

		//time.Sleep(time.Second)

		go func() {
			<-ch1
			fmt.Println("cat")
			ch2 <- struct{}{}
		}()

		//time.Sleep(time.Second)

		go func() {
			<-ch2
			fmt.Println("fish")
			ch3 <- struct{}{}
		}()

		//time.Sleep(time.Second)
		<-ch3
	}

}
