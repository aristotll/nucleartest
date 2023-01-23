package main

import (
	"arena"
	"flag"
	"fmt"
	"strconv"
	"time"
)
// gotip run -tags "goexperiment.arenas" main.go -arena 21
// GOEXPERIMENT=arenas gotip run main.go -arena 21
var n = 0
type Node struct {
	left, right *Node
	value       []byte
}
func bottomUpTree(depth int) *Node {
	if depth <= 0 {
		return &Node{}
	}
	return &Node{bottomUpTree(depth - 1), bottomUpTree(depth - 1), make([]byte, 128, 128)}
}
func bottomUpTreeWithArena(depth int, a *arena.Arena) *Node {
	node := arena.New[Node](a)
	node.value = arena.MakeSlice[byte](a, 128, 128)
	if depth <= 0 {
		return node
	}
	node.left = bottomUpTreeWithArena(depth-1, a)
	node.right = bottomUpTreeWithArena(depth-1, a)
	return node
}
func (n *Node) itemCheck() int {
	if n.left == nil {
		return 1
	}
	return 1 + n.left.itemCheck() + n.right.itemCheck()
}
const minDepth = 4
var useArena = flag.Bool("arena", false, "use arena")
func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		n, _ = strconv.Atoi(flag.Arg(0))
	}
	appStart := time.Now()
	defer func() {
		fmt.Printf("benchmark took: %v\n", time.Since(appStart))
	}()
	if *useArena {
		maxDepth := n
		if minDepth+2 > n {
			maxDepth = minDepth + 2
		}
		stretchDepth := maxDepth + 1
		a := arena.NewArena()
		start := time.Now()
		check := bottomUpTreeWithArena(stretchDepth, a).itemCheck()
		a.Free()
		fmt.Printf("stretch tree of depth %d\t check: %d, took: %v\n", stretchDepth, check, time.Since(start))
		a = arena.NewArena()
		longLiveStart := time.Now()
		longLivedTree := bottomUpTreeWithArena(maxDepth, a)
		defer a.Free()
		for depth := minDepth; depth <= maxDepth; depth += 2 {
			iterations := 1 << uint(maxDepth-depth+minDepth)
			check = 0
			start := time.Now()
			for i := 1; i <= iterations; i++ {
				a := arena.NewArena()
				check += bottomUpTreeWithArena(depth, a).itemCheck()
				a.Free()
			}
			fmt.Printf("%d\t trees of depth %d\t check: %d, took: %v\n", iterations, depth, check, time.Since(start))
		}
		fmt.Printf("long lived tree of depth %d\t check: %d, took: %v\n", maxDepth, longLivedTree.itemCheck(), time.Since(longLiveStart))
	} else {
		maxDepth := n
		if minDepth+2 > n {
			maxDepth = minDepth + 2
		}
		stretchDepth := maxDepth + 1
		start := time.Now()
		check := bottomUpTree(stretchDepth).itemCheck()
		fmt.Printf("stretch tree of depth %d\t check: %d, took: %v\n", stretchDepth, check, time.Since(start))
		longLiveStart := time.Now()
		longLivedTree := bottomUpTree(maxDepth)
		for depth := minDepth; depth <= maxDepth; depth += 2 {
			iterations := 1 << uint(maxDepth-depth+minDepth)
			check = 0
			start := time.Now()
			for i := 1; i <= iterations; i++ {
				check += bottomUpTree(depth).itemCheck()
			}
			fmt.Printf("%d\t trees of depth %d\t check: %d, took: %v\n", iterations, depth, check, time.Since(start))
		}
		fmt.Printf("long lived tree of depth %d\t check: %d, took: %v\n", maxDepth, longLivedTree.itemCheck(), time.Since(longLiveStart))
	}
}
