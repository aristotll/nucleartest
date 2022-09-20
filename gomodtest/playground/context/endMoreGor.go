package main

import (
	"context"
	"fmt"
	"time"
)


// 通过 context 控制多个 goroutine

type obj2 int64

func (o *obj2) handle(ctx context.Context, name ...string) {
	for {
		select {
		case <-ctx.Done():
			// 通过 Value 方法读取 ctx.Value(key)
			fmt.Println(name, "结束监控", ctx.Value("language"))
			return
		default:
			fmt.Println(name, "监控中...", ctx.Value("language"))
			<-time.Tick(time.Second)
		}
	}
}

func main() {
	var o obj2
	ctx, cancel := context.WithCancel(context.Background())
	// 通过 WithValue 传递元数据
	// 使用 context.WithValue 方法附加一对 K-V 的键值对，
	// 这里 Key 必须是等价性的，也就是具有可比性；Value 值要是线程安全的
	key := "language"
	value := context.WithValue(ctx, key, "go")

	// 启动 3 个 goroutine 进行监控，并将 context 传入
	go o.handle(ctx, "g1")
	go o.handle(ctx, "g2")
	go o.handle(ctx, "g3")

	go o.handle(value)
	go o.handle(value)
	go o.handle(value)
	time.Sleep(time.Second * 3)
	fmt.Println("通知停止监控")
	// 结束所有传入了 context 的 goroutine
	cancel()
	time.Sleep(time.Second * 3)
}
