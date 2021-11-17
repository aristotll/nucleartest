package main

import (
	"fmt"
	"time"
)

func main() {
	tf := time.AfterFunc(time.Second*5, func() {
		fmt.Println("timeout!")
	})

	go func() {
		select {
		case <-tf.C:
		}
	}()

	time.Sleep(time.Second * 10)
}
