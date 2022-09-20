package main

import "fmt"
import "time"

func test() {
	start := time.Now()
	l := make([]int, 0)
	for i := 0; i < 100000000; i++ {
		if i%10 == 0 {
			l = append(l, i)
		}
	}
	end := time.Since(start)
	fmt.Println("执行用时：", end)
}

func main() {
	//fmt.Println(111 / 29)
	test()
}
