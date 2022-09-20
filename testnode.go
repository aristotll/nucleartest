package main

import "fmt"

type TreeNode struct {
	Val int
}

func findLastSameNode(x, y []*TreeNode) (res *TreeNode) {
	xlast := len(x) - 1
	ylast := len(y) - 1
	i, j := 0, 0

	for i < xlast && j < ylast {
		if x[i].Val == y[j].Val {
			res = x[i]
		}
		i++
		j++
	}
	return
}

func findLastSameNode1(x, y []*TreeNode) (res *TreeNode) {
	i, j := 0, 0

    for i < len(x) && j < len(y) {
        if x[i] == y[j] {
            res = x[i]
        }
        i++
        j++
    }
    return
}

func main() {
	x := []*TreeNode{{3}, {5}}
	y := []*TreeNode{{3}, {1}}
	fmt.Println(findLastSameNode1(x, y))
}
