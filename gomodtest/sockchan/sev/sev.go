package main

import (
	"gomodtest/sockchan/global"
	"fmt"
)

func main() {
	fmt.Println("send data to global chan")
	global.Ch <- 1
	fmt.Println("send ok")
}