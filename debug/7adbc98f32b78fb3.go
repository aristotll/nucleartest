package main

import (
	"math"
)

func minWindow(s string, t string) (ret string) {
	l, r := 0, 0
	// 目标所需要的字母及次数，比如 t = "abcc"，那么该 map 就是 [a: 1, b: 1, c: 2]
	need := make(map[byte]int)
	// 当前窗口的字母及其出现次数
	window := make(map[byte]int)
	// 有几个字母匹配了，比如目标 string 是 abcc，当前窗口的元素是 [c, c, d]，有 2 个 c，与目标数量相同
	// 那么 c 就可以算匹配了
	valid := 0
	// 保存当前最小长度
	min := math.MaxInt64

	for _, c := range []byte(t) {
		need[c]++
	}

	for l <= r && r < len(s) {
		rv := s[r]
		// 更新当前窗口的状态
		window[rv]++
		// 如果当前窗口该字母数量与目标所需数量相同，valid++，表示已经有一个字母匹配要求
		if window[rv] == need[rv] {
			valid++
		}
		// 滑动右窗口
		r++
		// 当前窗口已经包含全部需要的字母，开始尝试缩小窗口，找到最优解
		// 这里判断条件是 len(need) 而不是 len(t)，因为存在这种情况：t = aa，那么 need = [a: 2]，window = [a: 2]，然后 valid 的值为 1
		// 如果判断条件式 len(t) == valid，则此时不满足条件，但是实际上此时已经匹配目标字符串了
		for len(need) == valid {
			// 如果当前匹配窗口的 string 长度比 min 还要小，则更新 min 并保存最新结果
			if r-l < min {
				ret = s[l:r]
				min = r - l
			}
			// lv 是被移出窗口的元素
			lv := s[l]
			// l++，表示将当前的 s[l] 移出窗口
			l++
			// 同样从当前窗口 map 中移除
			window[lv]--
			// 如果目标 string 中包含被移出的元素
			if _, ok := need[lv]; ok {
				// 移除后不满足目标 string 要求的出现次数
				// 比如当前窗口 [c, b, a, c]，目标字符串 "a, b, c"
				// 那么我们缩小右窗口，将最左边的 c 从窗口移出，那么窗口元素为 [b, a, c]
				// 此时的 c 为 1 个，但是任然 >= 目标字符串中 c 的数量，所以此时满足目标次数，
				// valid 就不用 - 1，同理，在相反的情况下，则 valid 需要 -1，比如目标字符串
				// 是 "a, c, c" 时
				if window[lv] < need[lv] {
					// 则匹配字母 - 1
					valid--
				}
			}
		}
	}
	if min == math.MaxInt64 {
		return ""
	}
	return
}

func main() {
	//minWindow("ADOBECODEBANC", "ABC")
	minWindow("aa", "aa")
}
