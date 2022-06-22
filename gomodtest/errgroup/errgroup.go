package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string
type Search func(ctx context.Context, query string) (Result, error)

func fakeSearch(kind string) Search {
	return func(_ context.Context, query string) (Result, error) {
		return Result(fmt.Sprintf("%s result for %q", kind, query)), nil
	}
}

func main_() {
	Google := func(ctx context.Context, query string) ([]Result, error) {
		g, ctx := errgroup.WithContext(ctx)

		searches := []Search{Web, Image, Video}
		results := make([]Result, len(searches))
		for i, search := range searches {
			i, search := i, search // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				result, err := search(ctx, query)
				if err == nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return nil, err
		}
		return results, nil
	}

	results, err := Google(context.Background(), "golang")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	for _, result := range results {
		fmt.Println(result)
	}
}

func main2() {
	group, ctx := errgroup.WithContext(context.Background())
	count := 10
	for i := 0; i < count; i++ {
		i := i // 闭包重新捕获变量 i
		group.Go(func() error {
			select {
			case <-ctx.Done():
				log.Printf("[%d] group 里有一个任务失败了，所以这个任务不会执行\n", i)
				return nil
			default:
			}
			if i == 3 {
				log.Printf("[%d] error\n", i)
				return fmt.Errorf("[%d] error", i)
			}
			log.Printf("[%d] success\n", i)
			return nil
		})
		time.Sleep(time.Millisecond * 10)
	}
	if err := group.Wait(); err != nil {
		log.Println(err)
	}
}

// 写的什么狗屁玩意
func main1() {
	// WithContext() 里面用 context.Cancel 包装了传入的 context，并返回包装后
	// 的 ctx，而另一个返回值 cancel 则是放到了 errgroup.Group 里面
	// 如果 g.Go(func() error) 里面的 func 返回了错误，那么会调用 Group.cancel
	// 将包装的 ctx 取消
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// 模拟单个服务错误退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	// g1
	// g1 退出了所有的协程都能退出么？
	// g1 退出后, context 将不再阻塞（这话第一眼没看懂，context 还会被阻塞？），g2, g3 都会随之退出
	// 修改版注释：g1 只有在 g2 调用 Shutdown() 后才会退出，同时返回一个 error
	// 之后 ctx 被 cancel，其他 goroutine 的 select 都会走 case <-ctx.Done() 分支，
	// 从而达到退出的目的（因为 select 没有加 default 分支，所以会被阻塞，直到有一个 case
	// 满足条件，所以上面说的 context 将不再阻塞，实际应该是 select 不再阻塞）
	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	g.Go(func() error {
		// g2 的 Shutdown 会让这里停止阻塞并返回一个 err
		// 之后 ctx 会被 cancel 掉，g3 也会停止
		if err := server.ListenAndServe(); err != nil {
			log.Println("g1 error: ", err)
			return err
		}
		return nil
	})

	// g2
	// g2 退出了所有的协程都能退出么？
	// g2 退出时，调用了 shutdown，g1 会退出
	// g2 退出后, context 将不再阻塞，g3 会随之退出
	// 修改：g2 退出后, g1 cancel 掉了 ctx，g3 的 select 将不再阻塞，g3 会随之退出
	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverOut: // curl localhost:8080/shutdown
			log.Println("server will out...")
		}

		cancelCtx, cancel := context.WithCancel(context.Background())
		// 这里不是必须的，但是如果使用 _ 的话静态扫描工具会报错，加上也无伤大雅
		defer cancel()

		log.Println("shutting down server...")
		// Shutdown() 会停止 g1，同时 g1 的 ListenAndServe() 会返回一个 error
		if err := server.Shutdown(cancelCtx); err != nil {
			log.Println("g2 error: ", err)
			return err
		}
		return nil
	})

	// g3
	// g3 捕获到 os 退出信号将会退出
	// g3 退出了所有的协程都能退出么？
	// g3 退出后, context 将不再阻塞，g2 会随之退出
	// 修改：g3 收到信号返回一个 error 并退出，ctx 会被 cancel 掉，然后 g2 的 select
	// 停止阻塞，执行下面的流程
	// g2 退出时，调用了 shutdown，g1 会退出
	// 然后 main 函数中的 g.Wait() 退出，所有协程都会退出
	g.Go(func() error {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("g3: ctx is done")
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}

func main() {
	// SetLimit 将这个组中的活动 goroutine 的数量限制为最多 n。负值表示没有限制。
	// 对 Go 方法的任何后续调用都将阻塞，直到它可以添加一个活动的 goroutine 而不会
	// 超过配置的限制。当组中的任何 goroutine 处于活动状态时，不得修改限制。
	eg := errgroup.Group{}
	count := 20
	mostRunning := 2
	eg.SetLimit(mostRunning)

	for i := 0; i < count; i++ {
		i := i
		eg.Go(func() error {
			log.Println(i)
			time.Sleep(300 * time.Millisecond)
			return nil
		})
	}

	eg.Wait()
}
