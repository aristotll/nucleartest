package main

import (
	"sync/atomic"
)

func main() {
	v := atomic.Value{}
	v.Load()
}
