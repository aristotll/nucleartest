package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func toStdin() {
	//defer wg.Done()
	fmt.Println("in stdin")
	for i := 0; i < 100; i++ {
		_, _ = io.WriteString(os.Stdin, strconv.Itoa(i)+"\n")
		time.Sleep(time.Second)
	}
}

func toStdout(ch chan struct{}) {
	//defer wg.Done()
	fmt.Println("in stdout")
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	ch <- struct{}{}
}

func main() {
	//wg.Add(2)
	ch := make(chan struct{})
	fmt.Println("in main")

	go toStdout(ch)
	go toStdin()
	<-ch
	//wg.Wait()
}
