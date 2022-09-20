package quickstart

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
)

var addr = "localhost:8080"

func ClientRun() {
	// 定义了使用什么方式来实现服务发现。
	// 在这里我们使用最简单的 Peer2PeerDiscovery（点对点），客户端直连服务器来获取服务地址
	// 点对点方式只能连接一台服务器
	c := client.NewPeer2PeerDiscovery("tcp@"+addr, "")

	// 创建了 XClient， 并且传进去了 FailMode、 SelectMode 和 默认选项。
	// FailMode 告诉客户端如何处理调用失败：重试、快速返回，或者 尝试另一台服务器。
	// SelectMode 告诉客户端如何在有多台服务器提供了同一服务的情况下选择服务器
	xClient := client.NewXClient("Arith", client.Failtry,
		client.RandomSelect, c, client.DefaultOption)
	defer xClient.Close()

	req := &Req{
		A: 10,
		B: 20,
	}

	rsp := &Rsp{}

	// 调用了远程服务并且同步获取结果
	err := xClient.Call(context.Background(), "Mul", req, rsp)
	if err != nil {
		log.Fatal("fail to call: ", err)
	}
	log.Printf("%d * %d = %d", req.A, req.B, rsp.C)
}

func ClientByAsync() {
	c := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	xClient := client.NewXClient("Arith", client.Failtry,
		client.RandomSelect, c, client.DefaultOption)

	req := &Req{
		A: 10,
		B: 20,
	}

	rsp := &Rsp{}
	ch := make(chan *client.Call, 1)

	call, err := xClient.
		Go(context.Background(), "Mul", req, rsp, ch)
	if err != nil {
		log.Fatal(err)
	}

	rspCall := <-call.Done
	if rspCall.Error != nil {
		log.Fatalf("fail to call by async: %v", rspCall.Error)
	} else {
		log.Printf("%d * %d = %d", req.A, req.B, rsp.C)
	}

}
