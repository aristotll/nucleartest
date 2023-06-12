package main

import "fmt"

var isAdd = make(map[string]struct{})

//var emptySli []string

// 思路：遍历 strs，拿到当前 str，然后遍历掉该 str 后面的所有 str，看两个 str 长度是否相同，且
// 字母是否相同（比如 abc 和 cab），如果一样则添加到数组作为结果，该轮遍历结束后添加到结果的二维数组
// 为了防止某个 str 被重复添加到结果二维数组，所以额外用一个 map 来保存 str 是否已被添加过
func groupAnagrams(strs []string) (res [][]string) {
	var emptyStrCount int
	for i := range strs {
		if strs[i] == "" {
			emptyStrCount++
			//emptySli = append(emptySli, "")
			continue
		}
		var sli []string
		// 每个 str 一定会出现在结果二维数组中，但是为了防止重复出现，需要使用 map 进行判断
		if _, ok := isAdd[strs[i]]; !ok {
			sli = append(sli, strs[i])
            isAdd[strs[i]] = struct{}{}
		}

		// 遍历后面的所有 str
		for j := i + 1; j < len(strs); j++ {
            fmt.Println(strs[i], strs[j], strs[i]==strs[j])
			// 长度不一样直接 continue
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			var m = make(map[rune]struct{})
			// 将 strs[i] 的每个字母添加到 map
			makeStrMap(strs[i], m)
			// 再看当前 strs[j] 是否匹配
			if isMatch(strs[j], m) {
                fmt.Println(isAdd)
				if _, ok := isAdd[strs[j]]; !ok {
					sli = append(sli, strs[j])
					isAdd[strs[j]] = struct{}{}
				}
			}
		}
		if len(sli) > 0 {
			res = append(res, sli)
		}
	}
	//fmt.Println(emptyStrCount)
	if emptyStrCount > 0 {
		genSli := func(n int) (s []string) {
			for i := 0; i < n; i++ {
				s = append(s, "")
			}
			return
		}
		res = append(res, genSli(emptyStrCount))
	}
	return res
}

func makeStrMap(str string, m map[rune]struct{}) {
	for _, c := range str {
		m[c] = struct{}{}
	}
}

func isMatch(str string, m map[rune]struct{}) bool {
	for _, c := range str {
		if _, ok := m[c]; !ok {
			return false
		}
	}
	return true
}

func main() {
	v := groupAnagrams([]string{"h", "h", "h"})
	fmt.Println(v)
}
