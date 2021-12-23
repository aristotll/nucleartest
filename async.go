package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Struct struct {
	X, Y, Z int64
}

func (s *Struct) XX() {
	defer wg.Done()
	s.X = 10
	time.Sleep(time.Second * 5)
}

func (s *Struct) YY() {
	defer wg.Done()
	s.Y = 20
	time.Sleep(time.Second * 3)
}

func (s *Struct) ZZ() {
	defer wg.Done()
	s.Z = 30
	time.Sleep(time.Second * 3)
}

func main() {
	var s Struct
	wg.Add(3)

	go s.XX()
	go s.YY()
	go s.ZZ()

	wg.Wait()

	fmt.Printf("%+v \n", s)
}
