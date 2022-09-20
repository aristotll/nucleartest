package main

import (
	"fmt"
	"gomodtest/sockchan/global"
	"time"
)

func main() {
	select {
	case <-time.After(time.Second * 8):
		fmt.Println("timeout !")
	case v := <-global.Ch:
		fmt.Println(v)
	}
}
