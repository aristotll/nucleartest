package main

func intToRoman(num int) string {
	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	c := []string{
		"M", "CM", "D", "CD", "C",
		"XC", "L", "XL", "X", "IX",
		"V", "IV", "I"}
	var ret string

	for num > 0 {
		for i := 0; i < len(n); i++ {
			if num >= n[i] {
				//fmt.Println("before: ", num)
				ret += c[i]
				num -= n[i]
				//fmt.Println("after: ", ret, num)
			}
		}
	}

	return ret
}

func main() {
	intToRoman(20)
}
