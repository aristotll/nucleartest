package main

import "fmt"

// Assignment to method receiver doesn't propagate to other calls

type List []int

func (l List) add() {
	//l = append(l, 123)
	l[0], l[1] = l[1], l[0]
	// main point: 0x119e9e0
	// func point: 0xc0000140b0
	fmt.Printf("func point: %p\n", l)
}

func (l *List) addP() {
	// main point &i: 0xc00000c060, i:0x119f9e0
	// func* point: 0xc00000c060, 0x119f9e0, 0xc00000e030
	fmt.Printf("func* point: %p, %p, %p\n", l, *l, &l)
	*l = append(*l, 123)
	// func* point: 0xc00000c060, 0xc0000140b8, 0xc00000e030
	fmt.Printf("func* point: %p, %p, %p\n", l, *l, &l)
}

func (l List) lens() int {
	return len(l)
}

func main() {
	var i List = make([]int, 0)
	i = append(i, 666, 555)
	// main point &i: 0xc00000c060, i:0x119f9e0
	fmt.Printf("main point &i: %p, i:%p\n", &i, i)
	i.add()
	l := i.lens()
	// []
	fmt.Println(i)
	fmt.Println(l)

	i.addP()
	// [123]
	fmt.Println(i)
	fmt.Println(l)
}
