package main

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	var l, r, valid int

	for _, c := range []byte(s1) {
		need[c]++
	}

	for l <= r && r < len(s2) {
		rv := s2[r]
		window[rv]++

		if window[rv] == need[rv] {
			valid++
		}

		if r-l+1 == len(s1) {
			if valid == len(need) {
				return true
			}
			lv := s2[l]
			l++

			if _, ok := need[lv]; ok {
				window[lv]--
				if window[lv] <= need[lv] && valid > 0 {
					valid--
				}
			}

		}

		r++
	}
	return false
}

func main() {
	//checkInclusion("ab", "eidboaoo")
	checkInclusion("adc", "dcda")
}
