package main

import (
	"net/rpc"
)

type Service struct {}
type Args {X, Y int64}
type Reply {R int64}

func (s *Service) XX(args *Args, reply *Reply) error {
	reply.R = args.X + args.Y 
}

func main() {
	rpc.
}
