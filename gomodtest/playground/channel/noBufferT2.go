package main

import "fmt"

//func main() {
//	ch := make(chan int)
//	go func() {
//		str := <- ch
//		fmt.Println(str)
//	}()
//	ch <- 3
//}

func main() {
	ch := make(chan int)
	go func() {
		ch <- 3
	}()
	fmt.Println(<- ch)
}
