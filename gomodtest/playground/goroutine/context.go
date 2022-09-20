package main

import (
	"context"
	"fmt"
	"time"
)

type Obj int

func (o *Obj) handle(c context.Context, duration time.Duration) {
	select {
	case <-c.Done():
		fmt.Println("handle", c.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	var obj Obj
	go obj.handle(ctx, time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}

}
