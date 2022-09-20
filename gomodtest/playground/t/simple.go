package t

// 一个简单的累加算法
func Sum(x int) int {
	r := 0
	for i := 0; i < x; i++ {
		r += i
	}
	return r
}
