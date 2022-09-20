package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var (
		a atomic.Pointer[int64]
		i = int64(100)
		j = int64(200)
	)

	a.Store(&i)
	a.CompareAndSwap(&i, &j)
	fmt.Println(*a.Load())

	var (
		da atomic.Pointer[time.Time]
	)
	now := time.Now()
	da.Store(&now)
	fmt.Println(*da.Load())
	time.Sleep(time.Second)
	now1 := time.Now()
	da.CompareAndSwap(&now, &now1)
	fmt.Println(*da.Load())

	var i64a atomic.Int64
	i64a.Store(1)
	i64a.Swap(2)
	fmt.Println(i64a.Load())
	if ok := i64a.CompareAndSwap(3, 1); !ok {
		fmt.Println("cas fail")
	}

	var (
		_ atomic.Int32
		_ atomic.Int64
		_ atomic.Uint64
		_ atomic.Uint32
		_ atomic.Bool
		_ atomic.Uintptr
	)
}
