package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill)

	c1 := make(chan os.Signal, 1)
	signal.Notify(c1, os.Interrupt)

	c2 := make(chan os.Signal, 1)
	signal.Notify(c2, syscall.SIGTERM)

	fmt.Printf("pid: %d \n", os.Getpid())

	go func() {
		<-c
		fmt.Println("recv SIGKILL!")
	}()

	go func() {
		<-c1
		fmt.Println("recv SIGINT!")
		os.Exit(0)
	}()

	go func() {
		<-c2
		fmt.Println("recv SIGTERM!")
		os.Exit(0)
	}()

	time.Sleep(time.Minute)
}
