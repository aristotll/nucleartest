package main

import "fmt"

func receiveNum(c chan int, num []int) {
	for _, v := range num{
		c <- v
	}
}

func printNum(c chan int) {
	// var num []int
	for v := range c{
		fmt.Println(v)
		if v == 0 {
			break
		}
	}
}

func main() {
	c := make(chan int)
	arr := []int{1,432,543,13,5756,76}
	go receiveNum(c, arr)
	printNum(c)
}

