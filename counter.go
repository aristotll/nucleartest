package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Val int64
}

func (c *Counter) GetAndIncrease() int64 {
	c.Val++
	return c.Val
}

func (c *Counter) GetAndIncreaseSafe() int64 {
	c.Lock()
	c.Val++
	c.Unlock()
	return c.Val
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{}
	count := 100
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			c.GetAndIncrease()
			//fmt.Println(ret)
		}()
	}
	wg.Wait()
	fmt.Println(c.Val)
	c.Val = 0

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			c.GetAndIncreaseSafe()
		}()
	}
	wg.Wait()
	fmt.Println(c.Val)

}
