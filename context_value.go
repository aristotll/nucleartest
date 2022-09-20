package main

import (
	"context"
	"fmt"
	"time"
)

func go1(ctx context.Context) {
	key := "name"
	val := ctx.Value(key)
	fmt.Printf("go1 get val: %v\n", val)

	ctx1 := context.WithValue(context.Background(), "name", "bbb")
	go go2(ctx1)
}

func go2(ctx context.Context) {
	key := "name"
	val := ctx.Value(key)
	fmt.Printf("go2 get val: %v\n", val)
}

func main() {
	ctx := context.WithValue(context.Background(), "name", "123")
	go go1(ctx)

	time.Sleep(time.Second * 5)
}
