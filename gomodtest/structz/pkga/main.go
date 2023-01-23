package pkga

import "fmt"

type Top struct {
	A   int64
	B   string
	XXX float64
}

func Do(t *Top) {
	fmt.Printf("%+v\n", t)
}