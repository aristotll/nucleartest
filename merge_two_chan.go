package main

func merge(c1 chan<- int, c2 chan<- int) (c3 chan int) {
	c3 = make(chan int)
	go func() {
		for {
			select {
			case v := <-c1:
				c3 <- v
			case v := <-c2:
				c3 <- v
			}
		}
	}()
	return
}
