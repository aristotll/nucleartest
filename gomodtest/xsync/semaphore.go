package xsync

import (
	"context"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

func init() {
	log.SetFlags(log.Ltime)
}

func Semaphore() {
	var (
		maxWorkers int64 = 5
		sem              = semaphore.NewWeighted(maxWorkers)
		out              = make([]int, 32)
		ctx              = context.Background()
	)

	for k := range out {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		log.Printf("num goroutine: %v, cur k: %v\n",
			runtime.NumGoroutine(), k)

		go func(i int) {
			defer sem.Release(1)
			out[i] = (i + 1) * 10
			time.Sleep(time.Second * 2)
			log.Printf("goroutine[%d] done. \n", i)
		}(k)
	}

	log.Println(out)
}
