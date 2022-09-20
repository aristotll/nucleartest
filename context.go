package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	for {
		fmt.Println("123")
		time.Sleep(time.Second)
	}
}
