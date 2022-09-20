package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Printf("%v timeout \n", time.Now())

	}()

	for {
		select {
		case v, ok := <-ctx.Done():
			fmt.Printf("%v timeout1 \n", time.Now())
			fmt.Printf("val: %v, closed: %v \n", v, ok)
		}
	}
}
