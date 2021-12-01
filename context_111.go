package main

import (
    "fmt"
    "context"
    "time"
)

func main() {
    ctx := context.Background()
    ctx1, cancel := context.WithTimeout(ctx, time.Second*3)
    defer cancel()

    go func(ctx context.Context) {
        select {
        case <-ctx.Done():
            fmt.Println("done")
        default:
            time.Sleep(time.Second * 5) // working...
            fmt.Println("goroutine1 done")
        }
    }(ctx1)


    time.Sleep(time.Second * 3)

    go func(ctx context.Context) {
        select {
        case <-ctx.Done():
            fmt.Println("done")
        default:
            time.Sleep(time.Second)
            fmt.Println("goroutine2 done")
        }
    }(ctx1)

    time.Sleep(time.Second * 100)
}
