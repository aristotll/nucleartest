package chan1

import (
	"fmt"
	//"chan1"
)

func Recv() {
	v := <-Ch
	fmt.Println("[chan2] receive data: ", v)
}
