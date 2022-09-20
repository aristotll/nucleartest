package main

import "fmt"

func t(ch chan int) {
	v := <- ch
	fmt.Println(v)
}

func main() {
	ch := make(chan int)

	go t(ch)
	ch <- 10
}
