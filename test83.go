package main

import (
	"container/list"
	"fmt"
	"strings"
	"unicode"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str string字符串
 * @return string字符串
 */
func compress(str string) string {
	// write code here
	mulstack := list.New() // 保存数字的栈
	letstack := list.New() // 保存字母的栈

	var (
		mul int
		res strings.Builder
	)

	for _, c := range str {
		if unicode.IsLetter(c) {
			res.WriteRune(c)
		} else if unicode.IsNumber(c) {
			mul = mul*10 + int(c-'0')
		} else if c == '[' {
			mulstack.PushBack(mul)
			letstack.PushBack(res.String())
			mul = 0
			res.Reset()
		} else if c == ']' {
			popmul := mulstack.Remove(mulstack.Back()).(int)
			var temp string
			for i := 0; i < popmul; i++ {
				temp += res.String()
			}

			popchar := letstack.Remove(letstack.Back()).(string)

			res.Reset()
			res.WriteString(popchar + temp)
			fmt.Println(res.String())
		}
	}

	return res.String()
}

func main() {
	s := compress("HG[3|B[2|CA]]F")
	fmt.Printf("s: %v\n", s)
}
