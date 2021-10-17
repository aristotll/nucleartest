package bloom

import (
	"fmt"
	"github.com/quasoft/bloomflt"
	"github.com/bits-and-blooms/bloom/v3"
	"strconv"
)

func bloom1() {
	filter := bloomflt.New(1000, 0.01)
	filter.AddString("123")
	if filter.ContainsString("123") {
		fmt.Println("ok")
	}

	for i := 1000; i < 2000; i++ {
		filter.AddUInt64(uint64(i))
	}
	for i := 0; i < 10000; i++ {
		if filter.ContainsUInt64(uint64(i)) {
			fmt.Printf("%v is ok \n", i)
		} else {
			fmt.Printf("%v is no \n", i)
		}
	}
}

func bloom2() {
	filter := bloom.New(3000, 4)

	for i := 1000; i < 2000; i++ {
		filter.AddString(strconv.Itoa(i))
	}
	for i := 0; i < 10000; i++ {
		if filter.TestString(strconv.Itoa(i)) {
			fmt.Printf("%v is ok \n", i)
		} else {
			fmt.Printf("%v is no \n", i)
		}
	}
}
