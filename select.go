package main

import (
	"context"
	"fmt"
	"time"
)

func f(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("f end")
			return
		}
	}()
	for {
		fmt.Println("123")
		time.Sleep(time.Second)
	}
}

func cont() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer func() {
		cancel()
	}()

}

func main() {
	cont()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("超时退出")
			return
		default:

		}
	}
}
