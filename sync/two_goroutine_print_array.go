package main

import (
	"fmt"
	"time"
)

func error1() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch := make(chan struct{})
	i := 0

	go func(i *int) {
		for *i < len(s) {
			if *i%2 == 1 {
				fmt.Printf("goroutine[1] => %v \n", s[*i])
			}
			*i += 1
			ch <- struct{}{}
		}
	}(&i)

	go func(i *int) {
		for *i < len(s) {
			<-ch
			if *i%2 == 0 {
				fmt.Printf("goroutinue[2] => %v \n", s[*i])
			}
			*i += 1
		}
	}(&i)

	time.Sleep(time.Second * 5)
}

func main() {
	//s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := []int{1123, 435, 1, 888, 999}
	ch := make(chan struct{})
	//i := 0

	go func() {
		for i := 0; i < len(s); i++ {
			if i%2 == 0 {
				fmt.Printf("goroutine[1] => %v \n", s[i])
			}
			ch <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < len(s); i++ {
			<-ch
			if i%2 == 1 {
				fmt.Printf("goroutinue[2] => %v \n", s[i])
			}
		}
	}()

	time.Sleep(time.Second * 5)
}
