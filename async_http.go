package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func HTTPGet(url string) chan []byte {
	fmt.Println("[HttpGet] http get1 start")
	done := make(chan []byte)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		fmt.Printf("[HttpGet] http get %v done! \n", url)
		done <- []byte{'a', 'b', 'c'}
		fmt.Println("[HttpGet] http get end")
	}()

	return done
}

func HandlerResponse(url string) {
	go func() {
		defer wg.Done()
		r := <-HTTPGet(url)
		fmt.Println("[HandlerResponse] get the resp from HttpGet")
		fmt.Println("resp: ", string(r))
		fmt.Println("[HandlerResponse] handler resp done!");
	}()
}

func DoOtherWork() {
	fmt.Println("[otherWork] do some thing....")
}

func main() {
	wg.Add(2)

	HandlerResponse("www.baidu.com")
	DoOtherWork()

	wg.Wait()

	// Output:
	// [otherWork] do some thing....
	// [HttpGet] http get1 start
	// [HttpGet] http get www.baidu.com done! 
	// [HttpGet] http get end
	// [HandlerResponse] get the resp from HttpGet
	// resp:  abc
	// [HandlerResponse] handler resp done!
}
