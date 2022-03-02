package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MySignal struct{}

func (m *MySignal) String() string {
	return "my signal"
}

func (m *MySignal) Signal() {
	fmt.Println("this is my signal!")
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, &MySignal{}, syscall.SIGUSR1)

	go func() {
		<-ch
		fmt.Println("SIGUSR1 is coming!")
	}()

	go func() {
		syscall.Kill(os.Getpid(), syscall.SIGUSR1)
	}()

	time.Sleep(time.Second * 5)
}
