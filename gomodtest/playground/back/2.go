package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type stu struct {
	Id int
	name string
	age int
	grade map[string]int
}

func t(ch chan *stu) {
	defer wg.Done()
	select {
	case s := <-ch:
		for subject, mark := range s.grade {
			fmt.Println(subject, mark)
		}
	default:
		fmt.Println("no data")
	}

}

func main() {
	wg.Add(1)
	ch := make(chan *stu)
	go t(ch)
	ch <- &stu{
		Id:   12,
		name: "zhang3",
		age:  123,
		grade: map[string]int{
			"chinese": 100,
			"math": 100,
		},
	}
	wg.Wait()

}
