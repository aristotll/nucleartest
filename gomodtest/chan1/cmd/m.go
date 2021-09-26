package main

import (
	gt "gomodtest/chan1"
	"time"
)

func main() {
	go gt.Send()
	go gt.Recv()

	time.Sleep(time.Second * 5)
}
