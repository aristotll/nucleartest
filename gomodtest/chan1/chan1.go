package chan1

import (
	"fmt"
)

var Ch = make(chan int64)

func Send() {
	fmt.Println("[ch1] send data to ch1!")
	Ch <- 1
}
