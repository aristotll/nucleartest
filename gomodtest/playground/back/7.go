package main

import "fmt"

type Tree struct {
	Val int
	Next *Tree
}

func te(i *int) {
	ii := i
	fmt.Printf("func i: %p, ii: %p \n", i, ii)
	*ii = 5
	//fmt.Println("参数 i :", i)
}

func tee(root *Tree) {
	cur := root
	cur = cur.Next
}

func main() {
	i := 123
	te(&i)
	fmt.Printf("main point i: %p, i value: %d\n", &i, i)

	t := &Tree{
		Val:  123,
		Next: &Tree{
			Val:  456,
			Next: nil,
		},
	}
	tee(t)
	fmt.Println(t)
}
