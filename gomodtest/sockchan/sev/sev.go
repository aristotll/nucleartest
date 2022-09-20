package main

import (
	"fmt"
	"gomodtest/sockchan/global"
)

func main() {
	fmt.Println("send data to global chan")
	global.Ch <- 1
	fmt.Println("send ok")
}
