package main

// defer 是在函数 return 后调用，还是在整个程序结束后调用？

import (
	"fmt"
	"time"
)

func f() {
	defer fmt.Println("1")
}

func f1() {
	defer fmt.Println("2")
}

func main() {
	f()
	time.Sleep(time.Second)
	f1()
	time.Sleep(time.Second)
}
