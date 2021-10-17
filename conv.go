package main

import (
	"container/list"
	"flag"
	"fmt"
	"unicode"
)

var input = flag.String("s", "", "输入一个中缀表达式")

// 中缀转后缀
func conv(s string) []rune {
	stack := list.New()
	var res []rune

	for _, c := range s {
		if unicode.IsDigit(rune(c)) {
			res = append(res, c) 
		}
		if c == '(' {
			stack.PushBack(c)
		}
		if c == ')' {
			for stack.Len() > 0 &&
				stack.Back().Value.(rune) != '(' {
				res = append(res, stack.Back().Value.(rune))
				stack.Remove(stack.Back())
			}
			stack.Remove(stack.Back()) // pop "("
		}
		if c == '+' || c == '-' {
			for stack.Len() > 0 &&
				stack.Back().Value.(rune) != '(' {
				pop := stack.Remove(stack.Back()).(rune)
				res = append(res, pop)
			}
			stack.PushBack(c)
		}
		if c == '*' || c == '/' {
			//peek := stack.Back().Value.(rune)
			for stack.Len() > 0 &&
				stack.Back().Value.(rune) != '(' &&
				stack.Back().Value.(rune) != '+' &&
				stack.Back().Value.(rune) != '-' {
				pop := stack.Remove(stack.Back()).(rune)
				res = append(res, pop)
			}
			stack.PushBack(c)
		}
	}

	for stack.Len() > 0 {
		pop := stack.Remove(stack.Back()).(rune)
		res = append(res, pop)
	}

	return res
}

func main() {
	flag.Parse()
	//s := conv("9 +   (3-1)*3+8/2")
	s := conv(*input)
	for _, v := range s {
		fmt.Printf("%v ", string(v))
	}
	fmt.Println()
}
