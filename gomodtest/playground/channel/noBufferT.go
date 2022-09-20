package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("in sum")
	c <- sum
}
func main() {
	ch := make(chan int)
	s := []int{1, 2, 3, 4, 5}
	// 放置在此会死锁，因为 chan 一直在等待接收者，且 chan 在 main 线程中，会造成主协程阻塞
	//
	// v := <- ch
	go sum(s, ch)
	v := <- ch
	fmt.Println(v)
}


