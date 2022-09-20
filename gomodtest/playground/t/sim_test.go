package t

import (
	"fmt"
	"testing"
)

func TestSumWrong(t *testing.T) {
	result := Sum(2067)
	if result == 210 {
		fmt.Println(true)
	}else {
		t.Errorf("wrong result where arg = 20")
	}
}

func TestSumTrue(t *testing.T) {
	result := Sum(21)
	if result == 210 {
		fmt.Println(true)
	}else {
		t.Errorf("wrong result where arg = 20")
	}
}
