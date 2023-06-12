package main

import "fmt"

func findAnagrams(s string, p string) (ret []int) {
	l, r := 0, 0
	window := make(map[byte]int)
	need := make(map[byte]int)
	valid := 0

	for _, c := range []byte(p) {
		need[c]++
	}

	for l <= r && r < len(s) {
		rv := s[r]

		if _, ok := need[rv]; ok {
			window[rv]++
			if window[rv] == need[rv] {
				valid++
			}
		}

		// 当前窗口的字符数量与 p 相同，可以开始比较是否是异位词了
		for r-l+1 >= len(p) {
			if valid == len(need) {
				ret = append(ret, l)
			}
			d := s[l]
			l++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}

		r++
	}
	return ret
}

func main() {
	//findAnagrams("cbeababacd", "abc")
	fmt.Println(findAnagrams("aabcb", "aabbc"))
}
