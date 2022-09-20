package quickstart

import (
	"context"
	"github.com/smallnest/rpcx/server"
)

// 请求
type Req struct {
	A int
	B int
}

// 响应
type Rsp struct {
	C int
}

type Arith int

func (a *Arith) Mul(ctx context.Context, req *Req, rsp *Rsp) error {
	rsp.C = req.A * req.B
	return nil
}

func ServerRun() {
	s := server.NewServer()
	// 注册服务，并命名为 Arith
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", ":8080")
}
