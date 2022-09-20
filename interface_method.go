package main

import (
	"fmt"
)

type A interface {
	Do()
}

func (a A) DoWhat() {
	a.Do()
}
