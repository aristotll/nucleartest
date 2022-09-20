package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 100 大王
// 101 小王
// 10 j
// 11 q
// 12 k
var allPoker = []int{
	100, 101,
	1, 1, 1, 1,
	2, 2, 2, 2,
	3, 3, 3, 3,
	4, 4, 4, 4,
	5, 5, 5, 5,
	6, 6, 6, 6,
	7, 7, 7, 7,
	8, 8, 8, 8,
	9, 9, 9, 9,
	10, 10, 10, 10,
	11, 11, 11, 11,
	12, 12, 12, 12,
}

// 农民1
type Farmer1 struct {
	Hand []int // 手牌
}

// 地主
type Landlord struct {
	Hand []int
}

// 农民2
type Farmer2 struct {
	Hand []int
}

func main() {
	f1 := new(Farmer1)
	f2 := new(Farmer2)
	l := new(Landlord)

	rand.Seed(time.Now().UnixNano())
	// 洗牌
	rand.Shuffle(len(allPoker), func(i, j int) {
		allPoker[i], allPoker[j] = allPoker[j], allPoker[i]
	})

	var i int
	// 留 3 张底牌
	for ; i < len(allPoker)-3; i++ {
		if i%3 == 0 {
			f1.Hand = append(f1.Hand, allPoker[i])
		} else if i%3 == 1 {
			f2.Hand = append(f2.Hand, allPoker[i])
		} else {
			l.Hand = append(l.Hand, allPoker[i])
		}
	}
	l.Hand = append(l.Hand, allPoker[i:]...)

	fmt.Println("farmer1: ", f1)
	fmt.Println("farmer2: ", f2)
	fmt.Println("landlord: ", l)
}
