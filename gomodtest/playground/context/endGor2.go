package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type obj1 int64
var Wg1 sync.WaitGroup

func (o *obj1) handle(ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("结束监控")
			return
		default:
			fmt.Println("监控中...")
			<-time.Tick(time.Second)
		}
	}
}

func main() {
	var o obj1
	// context.Background():
	// 返回一个空的 Context，这个空的 Context 一般用于整个 Context 树的根节点
	// context.WithCancel(parent): 创建一个可取消的子 Context
	ctx, cancel := context.WithCancel(context.Background())
	go o.handle(ctx)
	time.Sleep(time.Second * 3)
	// 发出取消指令，结束 goroutine
	cancel()
	time.Sleep(time.Second * 3)
}
