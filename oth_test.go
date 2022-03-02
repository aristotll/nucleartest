package design_patterns

import (
	"bufio"
	"bytes"
	"container/heap"
	"container/list"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
	"unicode"
	"unsafe"
)

// =============== other testgeneric =================
func TestChangeSlice(t *testing.T) {
	n := []int{1, 2, 3, 4, 5}
	fn := func(nums []int) {
		var nn []int
		for _, num := range nums {
			nn = append(nn, num+1)
		}
		n = nn
	}
	fn(n)
	fmt.Println(n)
}

//
//func Itob(num int) []byte {
//	ll := 0
//	nn := num
//	for nn > 0 {
//		ll++
//		nn /= 10
//	}
//
//	b := make([]byte, ll)
//	index := len(b) - 1
//	nn = num
//	for nn > 0 {
//		b[index] = byte(nn % 10)
//		index--
//		nn /= 10
//	}
//	fmt.Println(b)
//	return b
//}

func TestItob(t *testing.T) {
	i := 123
	fmt.Println(byte(i))
	Itob(123)
}

func compress(chars []byte) int {
	m := make(map[byte]int)

	for _, char := range chars {
		m[char]++
	}

	// 缩短数组
	//chars = chars[:len(m)]
	//fmt.Println(len(m))

	i := 0
	for char, num := range m {
		if num > 9 {
			chars[i] = char
			i++
			for _, n := range Itob(num) {
				chars[i] = n
				i++
			}
		} else if num == 1 {
			chars[i] = char
			i++
		} else {
			chars[i] = char
			i++
			chars[i] = byte(num)
			i++
		}
	}

	chars = chars[:i]
	//chars = c4

	for _, char := range chars {
		fmt.Printf("%v:%T ", string(char), string(char))
	}
	//fmt.Println(chars)
	return len(chars)
}

// 12 -> '1','2'
func Itob(num int) []byte {
	ll := 0
	nn := num
	for nn > 0 {
		ll++
		nn /= 10
	}

	b := make([]byte, ll)
	index := len(b) - 1
	nn = num
	for nn > 0 {
		b[index] = byte(nn % 10)
		index--
		nn /= 10
	}
	return b
}

func TestCompress(t *testing.T) {
	c := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	compress(c)
	a := 0
	b := '1' - '0'
	fmt.Println("==", a+int(b))
}

func addStrings(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	m := max(l1, l2)
	p1, p2 := l1-1, l2-1
	var sum float64
	flag := 0 // 进位
	i := 1

	for m > 0 {
		ss := 0
		if flag != 0 {
			ss += flag
			flag = 0
		}
		if p1 >= 0 {
			ss += int(num1[p1] - '0')
			p1--
		}
		if p2 >= 0 {
			ss += int(num2[p2] - '0')
			p2--
		}
		flag += ss / 10
		ss %= 10
		sum += float64(i * ss)
		i *= 10
		m--
	}
	fmt.Println(sum)
	return ""
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestStrAdd(t *testing.T) {
	addStrings("3876620623801494171", "6529364523802684779")
}

func addStrings1(num1 string, num2 string) string {
	l1, l2 := len(num1), len(num2)
	m := max(l1, l2)
	p1, p2 := l1-1, l2-1
	var sum []byte
	flag := 0 // 进位
	i := 1
	//var sb strings.Builder

	for m >= 0 {
		ss := 0
		if flag != 0 {
			ss += flag
			flag = 0
		}
		if p1 >= 0 {
			ss += int(num1[p1] - '0')
			p1--
		}
		if p2 >= 0 {
			ss += int(num2[p2] - '0')
			p2--
		}
		flag += ss / 10
		ss %= 10
		sum = append(sum, byte(ss+'0'))
		//sum += i * ss
		i *= 10
		m--
	}
	reverse(sum)
	//v := strconv.Itoa(sum)
	fmt.Println(string(sum), sum)
	return string(sum)
}

func reverse(b []byte) {
	l := len(b)
	for i, j := 0, l-1; i < j; i++ {
		b[i], b[j] = b[j], b[i]
		j--
	}
}

func TestStrAdd1(t *testing.T) {
	//addStrings1("3876620623801494171", "6529364523802684779")
	addStrings1("128", "476")
	//s := "12312312"
	//fmt.Println(int(s[0]))
}

func find4(n int) {
	s := make([]byte, n)
	memset(s)
	//fmt.Println(s)
	i := n - 1
	//temp := 0
	for i > 0 {
		if s[i] < '9' {
			s[i] += 1
		} else if s[i] == '9' && s[i-1] != '9' {
			s[i-1] += 1
			s[i] = '0'
		} else {
			s[i] = '0'
			s[i-1] = '0'
			i -= 2
		}
		fmt.Println(string(s))
	}
}

func memset(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = '0'
	}
}

func TestName(t *testing.T) {
	find4(5)
}

func backtrack(res *[][]int, output []int, first int, len int) {
	if first == len {
		//temp := make([]int, 0)
		//copy(temp, output)
		// append 可能会导致切片地址变更
		*res = append(*res, output)
		return
	}

	for i := first; i < len; i++ {
		swap(output, i, first)
		backtrack(res, output, first+1, len)
		swap(output, i, first)
	}
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(&res, nums, 0, len(nums))
	return res
}

func swap(n []int, x, y int) {
	n[x], n[y] = n[y], n[x]
}

func swapTest(n []int) {
	swap(n, 0, 2)
}

func Test111(t *testing.T) {
	nums := []int{1, 2, 3}
	r := permute(nums)
	fmt.Println(r)
}

func TestSwap(t *testing.T) {
	n := []int{1, 2, 3}
	//swap(n, 0, 2)
	swapTest(n)
	fmt.Println(n)
}

func permute1(nums []int) [][]int {
	// 1.初始化结果，用于存放结果
	var result [][]int
	// 2..判断临界条件，当为空切片时，返回空结果
	if len(nums) == 0 {
		return result
	}
	// 3.初始化其他变量
	// 创建中间变量，存放临时结果
	var temp []int
	// 创建bool值，判断该位置数字是否用过
	isvisited := make([]bool, len(nums))
	// 4.回溯函数
	BackTrack(isvisited, temp, nums, &result)
	return result
}

func BackTrack(isvisited []bool, temp []int, nums []int, result *[][]int) {
	// 5.判断回溯函数结束条件
	// 当临时切片长度和所给的数字长度相等时，将该切片加入结果
	if len(temp) == len(nums) {
		// 由于go语言的特性如果不特别说明创建的切片本质上都是指向同一个内存空间
		// 如果想要赋值的切片与原来切片不相关，需要另外开辟空间，这里用到copy函数，开辟独立空间
		current := make([]int, len(temp))
		copy(current, temp)

		*result = append(*result, current)
		return
	}

	// 6.遍历数组中的数字，进行排列组合
	for i := 0; i < len(nums); i++ {
		// 7.减枝，当该位置数字使用过时则跳过
		if isvisited[i] {
			continue
		}

		// 8.添加数字
		temp = append(temp, nums[i])
		// 将该位置数字设置为访问过的状态
		isvisited[i] = true
		// 9.继续搜索该支线
		BackTrack(isvisited, temp, nums, result)
		// 10.回溯，恢复到之前的状态
		temp = temp[:len(temp)-1]
		isvisited[i] = false
	}
}

func TestName2(t *testing.T) {
	n := []int{1, 2, 3}
	r := permute1(n)
	fmt.Println(r)
}

func TestBacktrackSpeed(t *testing.T) {
	n := []int{1, 2, 3}
	s := time.Now()
	r := permute(n)
	fmt.Println(r)
	fmt.Println(time.Since(s))

	s = time.Now()
	r1 := permute1(n)
	fmt.Println(r1)
	fmt.Println(time.Since(s))
}

// BenchmarkBool-8   	 2581156	       437.1 ns/op
// BenchmarkSwap-8   	 1999869	       573.2 ns/op
// BenchmarkSwap-8   	 2313684	       440.1 ns/op
func BenchmarkSwap(b *testing.B) {
	n := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		permute(n)
	}
}

// BenchmarkBool-8   	  999951	      1029 ns/op
// BenchmarkBool-8   	 1098356	       963.6 ns/op
// BenchmarkBool-8   	 1248417	       927.0 ns/op
func BenchmarkBool(b *testing.B) {
	n := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		permute1(n)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func TestFib(t *testing.T) {
	r := fib(4)
	fmt.Println(r)
}

func threeSum(nums []int) [][]int {
	start := time.Now()
	count := 0 // for 次数
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	k := 0
	i, j := k+1, len(nums)-1
	for k < len(nums)-2 {
		count++
		s := 0
		// [-1, 0, 0]
		if i >= j {
			k++
			continue
		}

		if nums[k] > 0 {
			break
		}
		//fmt.Println(k, i, j)
		s = nums[k] + nums[i] + nums[j]
		if s > 0 {
			for i < j && nums[j] == nums[j-1] {
				j--
			}
			j--
		} else if s < 0 {
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			i++
		} else {
			r := make([]int, 0)
			r = append(r, nums[k], nums[i], nums[j])
			res = append(res, r)
			for i < j && nums[i] == nums[i+1] {
				i++
			}
			for i < j && nums[j] == nums[j-1] {
				j--
			}
			i++
			j--

		}

		if i >= j {
			for k < len(nums)-1 && nums[k] == nums[k+1] {
				k++
			}
			k++
			i, j = k+1, len(nums)-1
		}
	}
	//fmt.Println(count)
	end := time.Since(start)
	fmt.Println("threeSum", end)
	return res
}

func threeSum1(nums []int) [][]int {
	start := time.Now()
	count := 0 // for 次数
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	res := make([][]int, 0)

	for k := 0; k < len(nums)-2; k++ {
		count++
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, len(nums)-1

		for i < j {
			count++
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
			} else if sum > 0 {
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			} else {
				r := make([]int, 0)
				r = append(r, nums[k], nums[i], nums[j])
				res = append(res, r)

				for i < j && nums[i] == nums[i+1] {
					i++
				}
				i++
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				j--
			}
		}
	}
	end := time.Since(start)
	fmt.Println("threeSum1", end)
	//fmt.Println(count)
	return res
}

func threeSum2(nums []int) [][]int {
	start := time.Now()
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	end := time.Since(start)
	fmt.Println(end)
	return ans
}

func TestThreeSum(t *testing.T) {
	//n := []int{-1, 0, 0}
	n1 := []int{-1, 0, 1, 2, -1, -4}
	//n2 := []int{-3, -3, 0, -5}
	//n3 := []int{0, 0, 0, 0, 0}
	sum := threeSum(n1)
	sum1 := threeSum1(n1)
	sum2 := threeSum2(n1)
	fmt.Println(sum)
	fmt.Println(sum1)
	fmt.Println(sum2)
}

func __append__(n *[]int) {
	fmt.Printf("[func point] %p\n", n)
	*n = append(*n, 123, 434, 13, 89)
	fmt.Printf("[func point] %p\n", n)
}

func __append(n []int) {
	fmt.Printf("[func] %p\n", n)
	n = append(n, 123, 434, 13, 89, 11, 1, 1, 1, 1, 1, 1, 11, 1, 11, 1, 1, 1, 1)
	fmt.Printf("[func] %p\n", n)
}

func TestName3(t *testing.T) {
	n := make([]int, 0)
	fmt.Printf("[main] %p\n", &n)
	__append__(&n)
	__append(n)
	fmt.Println(n)
}

func __int__(i *int) {
	fmt.Printf("[func] %p\n", i)
	*i = 5
}

func Test4(t *testing.T) {
	i := 123
	fmt.Printf("[main] %p\n", &i)
	__int__(&i)
}

// BenchmarkThreeSumBad-8   	 1243962	       999 ns/op
func BenchmarkThreeSumBad(b *testing.B) {
	n1 := []int{-1, 0, 1, 2, -1, -4}
	for i := 0; i < b.N; i++ {
		threeSum(n1)
	}
}

// BenchmarkThreeSumGood-8   	 3165510	       376 ns/op
func BenchmarkThreeSumGood(b *testing.B) {
	n1 := []int{-1, 0, 1, 2, -1, -4}
	for i := 0; i < b.N; i++ {
		threeSum1(n1)
	}
}

// BenchmarkThreeSumGood2-8   	 3429369	       309 ns/op
func BenchmarkThreeSumGood2(b *testing.B) {
	n1 := []int{-1, 0, 1, 2, -1, -4}
	for i := 0; i < b.N; i++ {
		threeSum2(n1)
	}
}

func TestTwoArraySort(t *testing.T) {
	var n [][]int
	n = append(n, []int{1, 9}, []int{2, 5}, []int{19, 20},
		[]int{10, 11}, []int{12, 20}, []int{0, 3},
		[]int{0, 1}, []int{0, 2})

	sort.Slice(n, func(i, j int) bool {
		return n[i][0] < n[j][0]
	})

	fmt.Println(n)
}

func reverseWords(s string) string {
	var sb strings.Builder
	ss := make([]string, 0)
	var word string

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if len(word) == 0 {
				continue
			} else {
				ss = append(ss, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	//ss = append(ss, word)
	if len(word) != 0 {
		ss = append(ss, word)
		word = ""
	}

	// if len(word) != 0 {

	//     word = ""
	// }

	for i, v := range ss {
		_reverse(&v)
		if i == len(ss)-1 {
			sb.WriteString(v)
		} else {
			sb.WriteString(v + ",")
		}
	}

	news := sb.String()
	//last := len(news) - 1
	//news = news[:len(news)-1]
	fmt.Println(news)

	// 去掉最后一个空格
	return news

}

func _reverse(s *string) {
	b := bytes.NewBufferString(*s).Bytes()

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	*s = string(b)
}

func TestReverseWords(t *testing.T) {
	s1 := "the sky is blue"
	s2 := "  hello world!  "
	s3 := "a good   example"

	reverseWords(s1)
	reverseWords(s2)
	reverseWords(s3)
}

func reverseString(s *string) {
	b := bytes.NewBufferString(*s).Bytes()

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	*s = string(b)
}

func TestReverseString(t *testing.T) {
	s := "12345"
	reverseString(&s)
	fmt.Println(s)
}

// bad
func reverseWords1(s string) string {
	ss := strings.TrimSpace(s)
	i, j := len(ss)-1, len(ss)-1
	flag := false
	var news strings.Builder

	for i >= 0 {
		for i >= 0 && ss[i] != ' ' {
			if !flag {
				flag = true
			}
			i--
			//fmt.Println(i)
		}
		_ = j

		if i < 0 || ss[i] == ' ' && flag {
			news.WriteString(ss[i+1 : j+1])
			news.WriteString(" ")
			flag = false
			i--
			j = i
		} else if ss[i] == ' ' && !flag {
			i--
			j--
		}
	}

	return strings.TrimSpace(news.String())
}

func reverseWord2(s string) string {
	j := len(s)
	i := j
	var sb strings.Builder

	for i >= 0 {
		for i >= 0 {
		}

	}

	return sb.String()
}

func _123() {
	fmt.Println("123")
}

func TestReverseWords1(t *testing.T) {
	//s1 := "the sky is blue"
	//s2 := "  hello world!  "
	s3 := "a good   example"
	//s4 := "aaa"

	var v string
	//v = reverseWords1(s1)
	//fmt.Println(v)
	//v = reverseWords1(s2)
	//fmt.Println(v)
	v = reverseWords1(s3)
	fmt.Println(v)
	//v = reverseWords1(s4)
	//fmt.Println(v)
	//_123()
}

// 找到根节点到某一节点的路径
//func findPath(node, need *TreeNode,
//	path, res *[]*TreeNode,
//	flag *bool) []*TreeNode {
//
//	if node == nil || *flag {
//		return *res
//	}
//	*path = append(*path, node)
//	if node == need {
//		*flag = true
//		*res = *path
//	}
//	findPath(node.Left, need, path, res, flag)
//	findPath(node.Right, need, path, res, flag)
//	*path = (*path)[:len(*path)-1]
//
//	rr := make([]*TreeNode, len(*res))
//	// 因为 slice 共用一个底层数组，第二次会更改第一次的结果，
//	// 所以需要 copy 一个新 slice
//	copy(rr, *res)
//
//	return rr
//
//}

func MoveToEnd(nums []int, len, start, end int) {
	last := len - 1
	for start <= end {
		nums[last] = nums[end]
		nums[end] = 0
		last--
		end--
	}
}

func TestMoveToEnd(t *testing.T) {
	n := []int{1, 2, 3, 0, 0, 0}
	MoveToEnd(n, len(n), 1, 2)
	fmt.Println(n)
}

func go1(nums []int, len, start, end int) {
	for end < len && start <= end {
		nums[end+1] = nums[end]
		nums[end] = 0
		end--
	}

}

func TestGo1(t *testing.T) {
	n := []int{1, 2, 3, 0, 0, 0}
	go1(n, len(n), 1, 2)
	fmt.Println(n)
}

func abs(x, y int) int {
	v := x - y
	if v < 0 {
		return -v
	}
	return v
}

func TestAbs(t *testing.T) {
	i := abs(51, 50)
	n := []int{1, 2, 3}
	nn := make([]int, len(n))
	copy(nn, n)
	fmt.Println(nn)
	fmt.Println(i)
}

func returnVal(n *[]int) []int {
	nn := make([]int, len(*n))
	s := time.Now()
	defer func() {
		fmt.Println(time.Since(s))
	}()
	copy(nn, *n)
	return nn
}

func returnPoint(n *[]int) *[]int {
	nn := make([]int, len(*n))
	s := time.Now()
	defer func() {
		fmt.Println(time.Since(s))
	}()
	copy(nn, *n)
	return &nn
}

func noReturn(n, nn *[]int) {
	s := time.Now()
	defer func() {
		fmt.Println(time.Since(s))
	}()
	copy(*nn, *n)
}

func TestReturn(t *testing.T) {
	l := 9999999
	n, nn := make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		n[i] = i
	}
	returnVal(&n)
	returnPoint(&n)
	noReturn(&n, &nn)
}

func backTrackByBool(nums []int, isVisited []bool, temp *[]int, res *[][]int) {
	if len(*temp) == len(nums) {
		newS := make([]int, len(*temp))
		copy(newS, *temp)
		*res = append(*res, newS)
	}

	for i := 0; i < len(nums); i++ {
		if isVisited[i] {
			continue
		}
		isVisited[i] = true
		*temp = append(*temp, nums[i])
		backTrackByBool(nums, isVisited, temp, res)
		isVisited[i] = false
		*temp = (*temp)[:len(*temp)-1]
	}
}

// BenchmarkBackTrackByBool-8   	  764626	      1324 ns/op
func BenchmarkBackTrackByBool(b *testing.B) {
	nums := []int{1, 2, 3}
	isVisited := make([]bool, len(nums))
	var temp []int
	var res [][]int
	for i := 0; i < b.N; i++ {
		backTrackByBool(nums, isVisited, &temp, &res)
	}
	//fmt.Println(res)
}

func backtrackBySwap(nums []int, res *[][]int, first int) {
	if first == len(nums) {
		newS := make([]int, len(nums))
		copy(newS, nums)
		*res = append(*res, newS)
	}

	for i := first; i < len(nums); i++ {
		swap(nums, i, first)
		backtrackBySwap(nums, res, first+1)
		swap(nums, i, first)
	}
}

// BenchmarkBacktrackBySwap-8   	  975050	      1209 ns/op
func BenchmarkBacktrackBySwap(b *testing.B) {
	nums := []int{1, 2, 3}
	//var temp []int
	var res [][]int
	for i := 0; i < b.N; i++ {
		backtrackBySwap(nums, &res, 0)
	}
}

// 26.14µs
func TestBacktrackBoolSpeed(t *testing.T) {
	start := time.Now()
	nums := []int{1, 2, 3}
	var temp []int
	var res [][]int
	isVisited := make([]bool, len(nums))
	backTrackByBool(nums, isVisited, &temp, &res)
	fmt.Println("bool: ", time.Since(start))

	//var temp1 []int
	//var res1 [][]int
	//start1 := time.Now()
	//backtrackBySwap(nums, &temp1, &res1, 0)
	//fmt.Println("swap: ", time.Since(start1))
}

// 26.603µs
func TestBacktrackSwapSpeed(t *testing.T) {
	//var temp1 []int
	var res1 [][]int
	nums := []int{1, 2, 3}
	start1 := time.Now()
	backtrackBySwap(nums, &res1, 0)
	fmt.Println("swap: ", time.Since(start1))
}

func TestBacktrackSwap(t *testing.T) {
	n := []int{1, 2, 3}
	//var temp []int
	var res [][]int
	backtrackBySwap(n, &res, 0)
	fmt.Println(res)
}

func compareVersion(version1 string, version2 string) int {
	p1, p2 := 0, 0

	for p1 < len(version1) || p2 < len(version2) {
		v1, v2 := 0, 0
		for p1 < len(version1) && version1[p1] != '.' {
			v, _ := strconv.Atoi(string(version1[p1]))
			v1 += v
			p1++
		}
		for p2 < len(version2) && version2[p2] != '.' {
			v, _ := strconv.Atoi(string(version2[p2]))
			v2 += v
			p2++
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
		p1++
		p2++
	}
	return 0
}

func TestCompareVersion(t *testing.T) {
	compareVersion("1.01", "1.001")
}

func swapStr(s *string, i, j int) {
	s1 := []byte(*s)
	s1[i], s1[j] = s1[j], s1[i]
	*s = string(s1)
}

func TestSwapStr(t *testing.T) {
	s := "123"
	swapStr(&s, 0, 2)
	fmt.Println(s)
}

func TestSliceSet(t *testing.T) {
	//n1 := []int{1, 2, 3}
	//n2 := []int{1, 2, 3}
	//n3 := []int{3, 4, 5}
	//n4 := []int{3, 4, 5}
	//n5 := []int{1, 2, 3}
	//var nn [][]int
	//nn = append(nn, n1, n2, n3, n4, n5)
	//setSlice2(&nn)
	//fmt.Println(nn)
	//
	//nn1 := [][]int{
	//	{1, 1, 1}, {1, 1, 1}, {1, 1, 1},
	//	{1, 1, 1}, {1, 1, 1}, {1, 1, 1},
	//}
	//setSlice2(&nn1)
	//fmt.Println(nn1)

	nn2 := [][]int{{2, 2, 1, 1}, {2, 2, 1, 1}, {2, 1, 2, 1}, {2, 1, 1, 2},
		{2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}, {2, 2, 1, 1}, {2, 1, 2, 1},
		{2, 1, 1, 2}, {2, 1, 1, 2}, {2, 1, 2, 1}, {1, 2, 2, 1}, {1, 2, 1, 2},
		{1, 2, 2, 1}, {1, 2, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2}, {1, 2, 1, 2},
		{1, 2, 2, 1}, {1, 1, 2, 2}, {1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}}
	setSlice2(&nn2)
	fmt.Println(nn2)

	//nn3 := [][]int{{2, 2, 1, 1}, {2, 2, 1, 1}}
	//setSlice2(&nn3)
	//fmt.Println(nn3)
}

func setSlice2(nn *[][]int) {
	for i := 0; i < len(*nn); i++ {
		for j := i + 1; j < len(*nn); j++ {
			if sliceEqual((*nn)[i], (*nn)[j]) {
				//*nn = append((*nn)[:j], (*nn)[j+1:]...)
				(*nn)[j] = []int{}
			}
		}
	}

	var news [][]int
	for _, v := range *nn {
		if len(v) != 0 {
			news = append(news, v)
		}
	}
	*nn = news

	//if len(*nn) > 1 {
	//	if sliceEqual((*nn)[0], (*nn)[1]) {
	//		*nn = append((*nn)[:1], (*nn)[2:]...)
	//	}
	//}
}

func sliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

type testNoInitMap struct {
	m map[int]int
}

func TestSliceEqual(t *testing.T) {
	n1 := []int{2, 2, 1, 1, 3}
	n2 := []int{2, 2, 1, 1, 2}
	fmt.Println(sliceEqual(n1, n2))
	fmt.Println(5 + 10*20/2)

	tm := testNoInitMap{}
	fmt.Println(tm.m[1])
}

func TestSliceSwap(t *testing.T) {
	n1 := []int{1, 3, 5, 6, 112}
	n2 := []int{12334, 234, 3424212}
	fmt.Printf("%p %p\n", n1, n2)
	n1, n2 = n2, n1
	fmt.Printf("%p %p\n", n1, n2)
}

var res int

func sumNums(n int) int {
	_ = n > 1 && sumNums(n-1) > 1
	res += n
	return res
}

func TestSumNums(t *testing.T) {
	r := sumNums(2)
	fmt.Println(r)
}

func __t__() {
	keys := []string{"MaxQueue", "max_value", "pop_front", "max_value", "push_back", "max_value", "pop_front", "max_value", "pop_front", "push_back", "pop_front", "pop_front", "pop_front", "push_back", "pop_front", "max_value", "pop_front", "max_value", "push_back", "push_back", "max_value", "push_back", "max_value", "max_value", "max_value", "push_back", "pop_front", "max_value", "push_back", "max_value", "max_value", "max_value", "pop_front", "push_back", "push_back", "push_back", "push_back", "pop_front", "pop_front", "max_value", "pop_front", "pop_front", "max_value", "push_back", "push_back", "pop_front", "push_back", "push_back", "push_back", "push_back", "pop_front", "max_value", "push_back", "max_value", "max_value", "pop_front", "max_value", "max_value", "max_value", "push_back", "pop_front", "push_back", "pop_front", "max_value", "max_value", "max_value", "push_back", "pop_front", "push_back", "push_back", "push_back", "pop_front", "max_value", "pop_front", "max_value", "max_value", "max_value", "pop_front", "push_back", "pop_front", "push_back", "push_back", "pop_front", "push_back", "pop_front", "push_back", "pop_front", "pop_front", "push_back", "pop_front", "pop_front", "pop_front", "push_back", "push_back", "max_value", "push_back", "pop_front", "push_back", "push_back", "pop_front"}
	values := []string{"", "", "", "", "46", "", "", "", "", "868", "", "", "", "525", "", "", "", "", "123", "646", "", "229", "", "", "", "871", "", "", "285", "", "", "", "", "45", "140", "837", "545", "", "", "", "", "", "", "561", "237", "", "633", "98", "806", "717", "", "", "186", "", "", "", "", "", "", "268", "", "29", "", "", "", "", "866", "", "239", "3", "850", "", "", "", "", "", "", "", "310", "", "674", "770", "", "525", "", "425", "", "", "720", "", "", "", "373", "411", "", "831", "", "765", "701", ""}
	output := []int{-999, -1, -1, -1, -999, 46, 46, -1, -1, -999, 868, -1, -1, -999, 525, -1, -1, -1, -999, -999, 646, -999, 646, 646, 646, -999, 123, 871, -999, 871, 871, 871, 646, -999, -999, -999, -999, 229, 871, 646, 285, 45, 646, -999, -999, 140, -999, -999, -999, -999, 837, 806, -999, 806, 806, 545, 806, 806, 806, -999, 561, -999, 237, 806, 806, 806, -999, 633, -999, -999, -999, 98, 866, 806, 866, 866, 866, 717, -999, 186, -999, -999, 268, -999, 29, -999, 866, 239, -999, 3, 850, 310, -999, -999, 806, -999, 674, -999, -999, 770}
	trueAns := []int{-999, -1, -1, -1, -999, 46, 46, -1, -1, -999, 868, -1, -1, -999, 525, -1, -1, -1, -999, -999, 646, -999, 646, 646, 646, -999, 123, 871, -999, 871, 871, 871, 646, -999, -999, -999, -999, 229, 871, 837, 285, 45, 837, -999, -999, 140, -999, -999, -999, -999, 837, 806, -999, 806, 806, 545, 806, 806, 806, -999, 561, -999, 237, 806, 806, 806, -999, 633, -999, -999, -999, 98, 866, 806, 866, 866, 866, 717, -999, 186, -999, -999, 268, -999, 29, -999, 866, 239, -999, 3, 850, 310, -999, -999, 770, -999, 674, -999, -999, 770}
	//fmt.Println(len(keys), len(values))
	//fmt.Println(keys[39])
	fmt.Printf("操作           操作值    输出    期望\n")
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%s     %s          %d     %d\n", keys[i], values[i], output[i], trueAns[i])
	}
}

func Test_(t *testing.T) {
	__t__()
}

func _tttt(i *int) *int {
	return i
}

func tt() {
	b := '1'
	s := "123"
	s = string(b) + s
	fmt.Println(s)
}

func Test_ttt(t *testing.T) {
	tt()
}

func decodeString(s string) string {
	stack := list.New()
	var res string
	var flag bool

	for _, char := range s {
		stack.PushBack(byte(char))
	}
	var cp string // 保存括号中的字母（也可能不是括号中的）
	for stack.Len() > 0 {
		top := stack.Back().Value.(byte)
		if top == '[' {
			if cp != "" && flag {
				res = cp + res
			}
			flag = true
		} else if top != ']' && flag == true {
			//fmt.Println("[vvv]:", int(top-'0'))
			for i := 0; i < int(top-'0'); i++ {
				res = cp + res
			}
			flag = false
			cp = ""
		} else if top != ']' {
			cp = string(top) + cp
		} else if top == ']' && cp != "" { // [cd]xyz
			res = cp + res
			cp = ""
		}
		stack.Remove(stack.Back())
		//fmt.Println(res)
	}
	if cp != "" {
		res = cp + res
	}

	return res
}

func TestDecodeString(t *testing.T) {
	println(decodeString("3[a2[c]]"))
}

func TestSliceAppend(t *testing.T) {
	b := make([]byte, 0, 512)
	for i := 0; i < 512; i++ {
		b = append(b, '1')
	}
	fmt.Println(cap(b))
	if len(b) == cap(b) {
		b = append(b, 0)[:len(b)]
	}
	fmt.Println(cap(b))
	s := 'a'
	ss := string(s)
	fmt.Println(ss)
}

func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go func(c context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("child goroutine1 cancel")
				return
			default:
				fmt.Println("goroutine1 running....")
				time.Sleep(time.Second)
				go func(ctx2 context.Context) {
					for {
						select {
						case <-ctx2.Done():
							fmt.Println("child child goroutine2 cancel")
							return
						default:
							fmt.Println("goroutine2 running....")
							time.Sleep(time.Second)
						}
					}
				}(c)
			}
		}
	}(ctx)
	time.Sleep(time.Second * 15)
}

func TestNN(t *testing.T) {
	n := []int{12, 13, 24, 66, 88, 123, 999}
	sort.Ints(n)
	i, j, max := 0, 1, math.MinInt32
	for i < len(n)-1 {
		sub := n[j] - n[i]
		if sub > max {
			max = sub
		}
		i++
		j++
	}
	fmt.Println(max)
}

func TestCopy(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)
}

func TestCopyHttp(t *testing.T) {
	http.HandleFunc("/t", func(w http.ResponseWriter, r *http.Request) {
		var bodyB []byte
		bodyB, _ = io.ReadAll(r.Body)
		fmt.Println(bodyB)
		//io.Copy(io.Discard, r.Body)
		r.Body = io.NopCloser(bytes.NewReader(bodyB))
		all, _ := io.ReadAll(r.Body)
		fmt.Println(all)

		all, _ = io.ReadAll(r.Body)
		fmt.Println(all)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}

}

func strToInt(str string) int {
	var res int
	var flag int // 1: 有负号 2: 正号
	str = strings.TrimSpace(str)

	for _, char := range str {
		// "9223372036854775808" int 无法容纳，需要在 for 中判断
		// "-91283472332"
		if res > math.MaxInt32 {
			if flag == 1 {
				res = math.MinInt32
				flag = 0
				break
			}
			res = math.MaxInt32
			break
		}
		if char == '-' { // 负号
			if flag == 2 { // 负号前面有正号了，例如 +-2
				break
			}
			flag = 1
		} else if char >= '0' && char <= '9' {
			res = res*10 + int(char-'0')
		} else if char == ' ' {
			break
		} else if char == '+' {
			if flag == 1 { // 正好前面有负号了，例如 -+1
				break
			}
			flag = 2
			continue
		} else { // a-z + ...
			break
		}
	}
	//fmt.Println(flag, res)
	if flag == 1 {
		res = -res
	}

	if res > math.MaxInt32 {
		res = math.MaxInt32
	} else if res < math.MinInt32 {
		res = math.MinInt32
	}
	return res
}

func TestAtoi(t *testing.T) {
	fmt.Println(strToInt("2147483648"))
	fmt.Println(strToInt("9223372036854775808"))
	fmt.Println(strToInt("-91283472332"))
	i, _ := strconv.Atoi("123 abc")
	fmt.Println(i)
	n := []int{1, 2, 3}
	n = append(n[:1], n[2:]...)
}

func atoi(t *testing.T) {
	state := "start"
	m := map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	_ = state
	_ = m
}

func Test415(t *testing.T) {
	s := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	http.HandleFunc("/t", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			log.Println(err)
			return
		}
		log.Println(s)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func TestStrR(t *testing.T) {
	s := "asdasd"
	for _, i := range s {
		fmt.Println(unicode.IsSpace(i))
		fmt.Println(i == '+')
	}
	var r rune
	fmt.Println(r == '1')
}

func Test401(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr)
		w.WriteHeader(http.StatusUnauthorized)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func TestUnicode(t *testing.T) {
	fmt.Println(unicode.Is(unicode.Han, '1'))
	l := list.New()

	n := 10
	for i := 0; i < n; i++ {
		l.PushBack(i)
	}
	l.PushBack(l.Front())

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func TestQ(t *testing.T) {
	n := []string{"1", "2", "3", "4", "5"}
	ns := fmt.Sprintf("%q", n)
	fmt.Println(ns)
}

func TestSet2(t *testing.T) {
	n := [][]int{{2, 2, 1, 1}, {2, 2, 1, 1}, {2, 1, 2, 1}, {2, 1, 1, 2},
		{2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}, {2, 2, 1, 1}, {2, 1, 2, 1},
		{2, 1, 1, 2}, {2, 1, 1, 2}, {2, 1, 2, 1}, {1, 2, 2, 1}, {1, 2, 1, 2},
		{1, 2, 2, 1}, {1, 2, 1, 2}, {1, 1, 2, 2}, {1, 1, 2, 2}, {1, 2, 1, 2},
		{1, 2, 2, 1}, {1, 1, 2, 2}, {1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}}

	var sf = func(n []int) string {
		return fmt.Sprintf("%q", n)
	}

	m := make(map[string]struct{})
	for _, v := range n {
		m[sf(v)] = struct{}{}
	}

	for s := range m {
		fmt.Printf("%s\n", s)
	}
}

func TestUnk(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	for v := l.Front().Value; v.(int) < 10; v = l.Front().Value {
		fmt.Println(v)
		l.Remove(l.Front())
	}
}

func TestFor(t *testing.T) {
	for i := 5; i < 10; {
		fmt.Println(i)
	}
}

func TestDefer(t *testing.T) {
	f := func() error {
		for i := 0; i < 5; i++ {
			if i == 2 {
				return errors.New("error")
			}
		}
		return nil
	}
	defer func() {
		fmt.Println("defer func")
	}()
	if err := f(); err != nil {
		fmt.Println(err)
		return
	}
	// 不会执行
	//defer func() {
	//	fmt.Println("defer func")
	//}()
}

func justReturnErr() error {
	return errors.New("a error")
}

func shadow() (err error) {
	if err := justReturnErr(); err != nil {

	}
	return nil
}

func shadows1() (i int) {
	if i := 5; i < 10 {

	}
	return 0
}

func Test7321(t *testing.T) {
	a, b := 0, 5
	m := (a + b) >> 1
	fmt.Println(m)

	n := []int{1, 2}
	n1 := n[1:1]
	fmt.Println(n1)

}

func TestSelectIO(t *testing.T) {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		fmt.Println("conn fd: ", conn)
	}

}

type aaaaa struct {
	a int
}
type bbbbb struct {
	b int
}

// 可否像 c 一样进行结构体强转？
func TestConvStruct(t *testing.T) {
	a := aaaaa{a: 123}
	b := bbbbb{b: 456}
	//a = b(a) // 并不可以
	_, _ = a, b
}

// 丑数
func isUgly(n int) bool {

	for i := 2; i < n-1; i++ {
		if i == 2 && n%i == 0 {
			continue
		} else if i == 3 && n%i == 0 {
			continue
		} else if i == 5 && n%i == 0 {
			continue
		} else if n%i == 0 {
			return false
		}
	}
	return true
}

func TestUglyNum(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v is ugly? %v\n", i, isUgly(i))
	}
}

func TestNetworkUnix(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		_ = conn
		conn.Write([]byte("conn server success"))
	}
}

func Test0xFFFF(t *testing.T) {
	fmt.Printf("%d\n", 0xFFFF)
}

func TestTcpRead(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 1)
		var bb []byte
		for {
			_, err := conn.Read(b)
			if err != nil {
				if err == io.EOF {
					log.Println("read EOF")
					break
				}
				log.Println("read error: ", err)
				break
			}
			bb = append(bb, b...)
			log.Println(string(b))
			fmt.Println("bb: ", string(bb))
		}
		fmt.Println("bb: ", string(bb))
	}
}

func TestSysSelect(t *testing.T) {
	// Create Socket
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln(err)
	}

	sockaddr := syscall.SockaddrInet4{}
	sockaddr.Addr = [4]byte{127, 0, 0, 1}
	sockaddr.Port = 8080
	// Bind
	if err := syscall.Bind(sockfd, &sockaddr); err != nil {
		log.Fatalln(err)
	}

	// Listen
	if err := syscall.Listen(sockfd, syscall.SOMAXCONN); err != nil {
		log.Fatalln(err)
	}

	for {
		// Accept
		connfd, sa, err := syscall.Accept(sockfd)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("sockaddr: %+v", sa)
		log.Println("connfd: ", connfd)

		var rset syscall.FdSet
		rset.Bits[int32(syscall.Stdout)] = 1
		rset.Bits[connfd] = 1

		log.Println("in select...")
		t := &syscall.Timeval{}
		t.Sec = 20
		t.Usec = 0
		if err := syscall.Select(connfd+1, &rset, nil, nil, t); err != nil {
			log.Fatalln(err)
		}
		log.Println("select over")
		b := make([]byte, 4096)
		for fd, ready := range rset.Bits {
			if ready == 1 {
				log.Printf("%d is ready\n", fd)
				for {
					_, err := syscall.Read(fd, b)
					if err != nil {
						if err == io.EOF {
							log.Println("read EOF")
							break
						}
						log.Println("read error: ", err)
						break
					}
					log.Println(b)
				}
			} else {
				log.Printf("%d is not ready\n", fd)
			}
		}
	}

	//_, _ = syscall.Write(connfd, []byte("testgeneric"))
}

// blog.csdn.net/gophers/article/details/33313959
func TestSysSocket(t *testing.T) {
	// Create Socket
	sockfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln(err)
	}

	sockaddr := syscall.SockaddrInet4{}
	sockaddr.Addr = [4]byte{127, 0, 0, 1}
	sockaddr.Port = 8080
	// Bind
	if err := syscall.Bind(sockfd, &sockaddr); err != nil {
		log.Fatalln(err)
	}

	// Listen
	if err := syscall.Listen(sockfd, syscall.SOMAXCONN); err != nil {
		log.Fatalln(err)
	}

	for {
		// Accept
		connfd, sa, err := syscall.Accept(sockfd)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("sockaddr: %+v", sa)
		_, _ = syscall.Write(connfd, []byte("testgeneric"))
	}
}

func TestByteSliceConvStr(t *testing.T) {
	b := []byte{'1', '2', '3', '4'}
	var s, s1 string

	start := time.Now()
	s = string(b)
	log.Println(s)
	u := time.Since(start)
	log.Println(u)

	start = time.Now()
	s1 = *(*string)(unsafe.Pointer(&b))
	log.Println(s1)
	u = time.Since(start)
	log.Println(u)
}

func TestIdiSort(t *testing.T) {
	n := []int{3, 30, 34, 5, 9}
	ns := make([]string, len(n))
	maxl := math.MinInt64
	for i := 0; i < len(n); i++ {
		s := strconv.Itoa(n[i])
		if len(s) > maxl {
			maxl = len(s)
		}
	}
	fmt.Println("maxl: ", maxl)
	for i := 0; i < maxl; i++ {
		sort.Slice(ns, func(x, y int) bool {
			if len(ns[x])-1 <= i && len(ns[y])-1 > i {
				if ns[y][i] == 0 {
					return true
				}
				return false
			} else if len(ns[y])-1 <= i && len(ns[x])-1 > i {
				if ns[x][i] == 0 {
					return true
				}
				return false
			}
			return ns[x][i] < ns[y][i]
		})
	}

	fmt.Println(ns)
}

func TestSSSS(t *testing.T) {
	n := []int{3, 30, 34, 5, 9}
	// 1 0
	// 2 1
	// 3 2
	// 4 3
	sort.Slice(n, func(i, j int) bool {
		fmt.Println(i, j)
		return false
	})
	i, _ := strconv.Atoi("   -1")
	fmt.Println(i)
}

func changeArray2(arr []int) []int {
	p1, p2 := 0, len(arr)-1
	for p1 < p2 /*&& p1 < len(arr) && p2 < len(arr)*/ {
		if arr[p1]%2 == 0 && arr[p2]%2 != 0 {
			arr[p1], arr[p2] = arr[p2], arr[p1]
			p1++
			p2--
			//fmt.Println("in swap")
		} else if arr[p1+1]%2 == 0 {
			p1++
		} else if arr[p2-1]%2 != 0 {
			p2--
		}
	}
	//fmt.Println(arr)
	return arr
}

func changeArray3(nums []int) []int {
	p1, p2 := 0, 0
	for p1 <= p2 {
		if nums[p1]%2 == 0 && nums[p2]%2 != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			if p1 < len(nums) {
				p1++
			}
		} else {
			if p2 < len(nums) {
				p2++
			}
		}
	}
	return nums
}

// 首位指针
func changeArray4(nums []int) []int {
	p1, p2 := 0, len(nums)-1
	for p1 < p2 {
		// p1 定位到奇数
		for p1 < len(nums)-1 && nums[p1]%2 != 0 {
			p1++
		}
		// p2 定位到偶数
		for p1 < p2 && nums[p2]%2 == 0 {
			p2--
		}
		// 交换
		nums[p1], nums[p2] = nums[p2], nums[p1]
		// 交换后更新 p1, p2
		p1++
		p2--
	}
	return nums
}

// 快慢指针
func changeArray5(nums []int) []int {
	now := time.Now()
	f, l := 0, 0

	for l < len(nums) {
		// 更新 f 的位置
		f = l
		// l 当前是奇数，则向前移动
		if nums[l]%2 != 0 {
			l++
			continue
		}
		// f 不断向前直到为奇数
		for f < len(nums)-1 && nums[f]%2 == 0 {
			f++
		}
		if nums[l]%2 == 0 && nums[f]%2 != 0 {
			// 交换 l, f
			nums[l], nums[f] = nums[f], nums[l]
		}
		// 交换后 l++
		l++
	}
	use := time.Since(now)
	fmt.Println("bad: ", use)
	return nums
}

// 快慢指针优化版
func changeArray55(nums []int) []int {
	now := time.Now()
	f, l := 0, 0

	for f < len(nums) {
		if nums[f]%2 != 0 {
			nums[f], nums[l] = nums[l], nums[f]
			l++
		}
		f++
	}
	use := time.Since(now)
	fmt.Println("good: ", use)
	return nums
}

func Test1675(t *testing.T) {
	//a := []int{2, 4, 1, 8, 4, 3, 5, 1, 123, 5}
	//a := []int{1, 2, 3, 4}
	a := []int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1}
	//a := []int{}
	//a := []int{1}
	//a := []int{1, 3, 5, 2, 4, 6}
	//a := []int{2, 4, 6, 1, 3, 5}
	//a := []int{1, 3, 5}
	//a := []int{2, 4, 6}
	r := changeArray5(a)
	_ = changeArray55(a)
	fmt.Println(r)
}

func Test_Upload_Download(t *testing.T) {
	mux := http.DefaultServeMux
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		file, header, err := r.FormFile("file")
		if err != nil {
			io.Copy(w, strings.NewReader("上传失败，未选择文件"))
			log.Println("get file error: ", err)
			return
		}
		path := "/Users/zz/Downloads/cp10/demo/src/main/java/com/example/demo/file/upload/"
		newfile, err := os.OpenFile(path+header.Filename,
			os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Println("create new file error: ", err)
			return
		}
		_, err = io.Copy(newfile, file)
		if err != nil {
			log.Println("io.copy error: ", err)
			return
		}
		io.Copy(w, strings.NewReader("文件上传成功"))
	})

	mux.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("content-type", "application/octet-stream")

		rootDir := "/Users/zz/Downloads/"
		if err := r.ParseForm(); err != nil {
			log.Println("parse form error: ", err)
			return
		}
		filename := r.Form.Get("filename")
		log.Println(filename, r.Form)
		log.Println(rootDir + filename)
		file, err := os.Open(rootDir + filename)
		if err != nil {
			log.Println("read file error: ", err)
			return
		}
		w.Header().Set("content-disposition", "attachment; filename=\""+filename+"\"")

		_, err = io.Copy(w, file)
		if err != nil {
			log.Println("read file error: ", err)
			return
		}
	})

	mux.HandleFunc("/upload/more", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		r.ParseMultipartForm(32 << 20)
		fs := r.MultipartForm.File["file"]
		l := len(fs)
		for i := 0; i < l; i++ {
			file, err := fs[i].Open()
			if err != nil {
				io.Copy(w, strings.NewReader("上传失败，未选择文件"))
				log.Println("get file error: ", err)
				return
			}
			path := "/Users/zz/Downloads/cp10/demo/src/main/java/com/example/demo/file/upload/"
			newfile, err := os.OpenFile(path+fs[i].Filename,
				os.O_RDWR|os.O_CREATE, 0777)
			if err != nil {
				log.Println("create new file error: ", err)
				return
			}
			_, err = io.Copy(newfile, file)
			if err != nil {
				log.Println("io.copy error: ", err)
				return
			}
			io.Copy(w, strings.NewReader("文件上传成功"))
		}
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func TestRetroreFlectServer(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	rok, wok := make(chan struct{}), make(chan struct{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		b := make([]byte, 4096)
		go func() {
			for {
				_, err = conn.Read(b)
				if err != nil {
					if err == io.EOF {
						log.Println("read EOF!")
						break
					}
					io.Copy(conn, strings.NewReader("read data from conn error: "+err.Error()))
					log.Println("read error: ", err)
					break
				}
			}

			rok <- struct{}{}
		}()

		go func() {
			l := len(b)
			for len(b) > 0 {
				n, err := conn.Write(b)
				l -= n
				if err != nil {
					log.Println("write error: ", err)
					io.Copy(conn, strings.NewReader("write data to conn error: "+err.Error()))
					break
				}
				wok <- struct{}{}
			}
		}()

		select {
		case <-rok:
		case <-wok:

		}
	}
}

func TestRetroreFlectClient(t *testing.T) {
	dial, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(dial, os.Stdin)
	all, err := io.ReadAll(dial)
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(os.Stdout, bytes.NewReader(all))
}

func TestInt32(t *testing.T) {
	fmt.Println(math.MinInt32, math.MaxInt32)
}

func Test1(t *testing.T) {
	c := make(chan struct{})
	n := 10
	for i := 0; i < n; i++ {
		i := i
		go func() {
			fmt.Println(i)
			c <- struct{}{}
		}()
	}
	for i := 0; i < n; i++ {
		<-c
	}
}

func TestName111(t *testing.T) {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		lis.Accept()
	}
}

func TestCString(t *testing.T) {
	//var str string
	//str += nil
	//var sb strings.Builder
	//sb.WriteString(nil)
	//fmt.Println(str)
}

func TestIp(t *testing.T) {
	var b [4]byte
	b = [4]byte{127, 0, 0, 1}
	fmt.Println(b)

	var bb [4]byte
	bb = [4]byte{128, 'b', 'c', 'd'}
	fmt.Println(string(bb[0]))
}

func findContinuousSequence(target int) [][]int {
	var res [][]int
	n := make([]int, (target>>1)+1)
	for i := 0; i < len(n); i++ {
		n[i] = i + 1
	}
	//fmt.Println(n)

	i, j := 0, 1
	sum := n[i] + n[j]
	for i < j {
		if sum < target {
			j++
			sum += n[j]
		} else if sum > target {
			sum -= n[i]
			i++
		} else if sum == target {
			var temp []int
			for i1 := i; i1 <= j; i1++ {
				temp = append(temp, n[i1])
			}
			res = append(res, temp)
			sum -= n[i]
			i++
			if j < len(n)-1 {
				j++
				sum += n[j]
			}
		}
	}
	return res
}

func TestOffer57(t *testing.T) {
	sequence := findContinuousSequence(15)
	fmt.Println(sequence)
}

func isStraight(nums []int) bool {
	sort.Ints(nums)
	var totalZero int // 有多少个 0
	for nums[totalZero] == 0 {
		totalZero++
	}
	curZero := totalZero

	for i := len(nums) - 1; i > totalZero; i-- {
		b := nums[i] - nums[i-1]
		if b == 1 {
			continue
		} else if b == 0 {
			return false
		} else if b > 1 {
			if b-1 > curZero {
				return false
			} else {
				curZero -= b - 1
			}
		}
	}
	return true
}

func TestOffer61(t *testing.T) {
	isStraight([]int{0, 0, 1, 2, 5})
}

func TestSliceNil(t *testing.T) {
	n := []int{}
	fmt.Println(len(n) == 0)
	fmt.Println(n == nil)
}

func Test6123(t *testing.T) {
	fmt.Println(4 ^ 1)
	fmt.Println(5 ^ 1)
	fmt.Println(5 ^ 2)
	fmt.Println(4 ^ 1 ^ 4 ^ 6)
	fmt.Println(1 ^ 2 ^ 10 ^ 4 ^ 1 ^ 4 ^ 3 ^ 3)

	ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := (ns[0] + ns[len(ns)-1]) * len(ns) / 2
	fmt.Println(res)

	fmt.Println(1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9)
}

func TestConvInt(t *testing.T) {
	var i int32 = 10
	var ii int
	s := time.Now()
	ii = int(i)
	e := time.Since(s)
	fmt.Println(e)

	s = time.Now()
	ii = *(*int)(unsafe.Pointer(&i))
	e = time.Since(s)
	fmt.Println(e)

	_ = ii
}

func TestTcpHttp(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	b := make([]byte, 4096)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		conn.Read(b)
		log.Println(string(b))
	}
}

func TestListenHttp(t *testing.T) {
	listener, err := net.Listen("tcp", ":http")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(conn.RemoteAddr())
	}
	//http.ListenAndServe()
	//http.CookieJar()
}

func TestFloat(t *testing.T) {
	var res float64 = 15 + 25 + 5.2
	fmt.Println(res)
}

func TestTempPort(t *testing.T) {
	listen, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		addr := listen.Addr().String()
		log.Println(addr)
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		conn.Write([]byte("123"))
	}
}

func TestForward(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 4096)
		buf := bytes.NewBuffer(b)
		n, err := io.Copy(buf, conn)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(n, buf.String())

		conn.Close()
	}
}

func TestTCP_Heartbeat(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listen.Accept()
		log.Println(conn.RemoteAddr())
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 1024)
		for {
			conn.Read(b)
			fmt.Println(string(b))
		}
		//log.Println("timeout, this conn will be close.")
	}
}

func TestCompareStr(t *testing.T) {
	s1 := "123"
	s2 := "456"
	log.Println(s1 > s2)
}

func Test8080(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	http.ListenAndServe(":8080", nil)
}

func Test80801(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		conn.RemoteAddr()
	}
}

func Test8080dial(t *testing.T) {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	conn.RemoteAddr()
	conn.Write([]byte("123"))

	select {}
}

func TestMaxInt(t *testing.T) {
	i := 1688849860263935
	log.Println(i)
	log.Println(math.MaxInt64)
}

func spiralOrder(matrix [][]int) []int {
	var (
		ylen  = len(matrix)
		xlen  = len(matrix[0])
		up    = 0
		down  = ylen - 1
		left  = 0
		right = xlen - 1
		res   []int
	)

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		fmt.Println("上边界：", res)
		up++
		if up > ylen {
			break
		}

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		fmt.Println("右边界：", res)
		right--
		if right < 0 {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		fmt.Println("下边界：", res)
		down--
		if down < 0 {
			break
		}

		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		fmt.Println("左边界：", res)
		left++
		if left > xlen {
			break
		}

		//fmt.Println(res)
	}

	return res
}

func TestTopK_inSlice(t *testing.T) {
	s := []int{1, 1, 2, 2, 3, 1, 3, 3, 3, 5, 4, 4, 4, 4, 5, 5}
	var s1 []int
	m := make(map[int]int)
	for _, ss := range s {
		if _, ok := m[ss]; !ok {
			s1 = append(s1, ss)
		}
		m[ss]++
	}
	fmt.Println(s1)
	fmt.Println(m)
	sort.Slice(s1, func(i, j int) bool {
		return m[s1[i]] > m[s1[j]]
	})

	fmt.Println(s1)
}

func Test517(t *testing.T) {
	s := []int{1, 2, 3}
	p := &s
	r := *p
	r[0] = 999
	fmt.Println(s[0])
}

func Test1Offer56(t *testing.T) {
	n := []int{4, 1, 4, 6}
	var j int

	for _, i := range n {
		log.Printf("%d: binary: %b\n", i, i)
		j ^= i
	}

	log.Printf("%d: binary: %b\n", j, j)
	// 1 0 0	4
	// 0 0 1	1
	// 1 1 0	6

	// 1 1 1	7

	log.Println((6 >> 0) & 1) // 0
	log.Println((6 >> 1) & 1) // 1
	log.Println((6 >> 2) & 1) // 1
	log.Println((6 >> 3) & 1) // 0
	log.Println((6 >> 4) & 1)

	log.Printf("%b\n", 10000)

	var n1 []int
	var n2 []int
	for i := 0; i < 16; i++ {
		b := (j >> i) & 1
		for ii := 0; ii < len(n); ii++ {
			nb := (n[ii] >> i) & 1
			if nb == b {
				n1 = append(n1, n[ii])
			} else {
				n2 = append(n2, n[ii])
			}
			if len(n1)+len(n2) == len(n) {
				log.Println(n1, n2)
				goto loop
			}
		}
	}

loop:
	var r1, r2 int
	for _, i := range n1 {
		r1 ^= i
	}

	for _, i := range n2 {
		r2 ^= i
	}

	log.Println(r1, r2)
	//log.Println(n1, n2)
}

func Test2Offer56(t *testing.T) {
	n := []int{1, 2, 10, 4, 1, 4, 3, 3}
	var j int

	for _, i := range n {
		j ^= i
	}

	d := 1
	// 找到 j 中第一个为 1 的位
	for d&j == 0 {
		d <<= 1
	}

	log.Println(d)
	var x, y int
	for _, i := range n {
		if d&i == 1 {
			x ^= i
		} else {
			y ^= i
		}
	}

	log.Println(x, y)
}

func TestBinary(t *testing.T) {
	log.Printf("%b\n", 3)
}

func TestPosixListen(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalln(err)
	}

	if err := syscall.Connect(fd, &syscall.SockaddrInet4{
		Port: 0,
		Addr: [4]byte{},
	}); err != nil {
		return
	}

	if err := syscall.Bind(fd, &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}); err != nil {
		log.Fatalln(err)
	}

	if err := syscall.Listen(fd, 1); err != nil {
		log.Fatalln(err)
	}

	for {
		connfd, _, err := syscall.Accept(fd)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(connfd)
	}
}

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	var p0, p1 int
	var i int

	for ; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
		//fmt.Println(p0, p1)
	}

	//fmt.Println(i, p0, p1)
	i--

	for nums[i] != 2 {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}

}

func TestLC75(t *testing.T) {
	n := []int{2, 0, 2, 1, 1, 0}
	sortColors(n)
	//listen, err := net.Listen("")
	//conn, err := listen.Accept()
	//conn1, err := net.Dial()
	//conn1.
}

func TestTCP1(t *testing.T) {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 1024)
		conn.Read(b)
		log.Println(string(b))
	}
}

func Test1510(t *testing.T) {
	fmt.Println(15 % 10)
}

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 1
		}
		return 0
	}

	if len(nums) == 0 {
		return 0
	}

	minlen := math.MaxInt64
	i, j := 0, 1
	sum := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	sum += nums[i] + nums[j]
	//cout << sum << endl;

	for i < j && j < len(nums) {
		if sum > target {
			minlen = min(minlen, j-i+1)
			sum -= nums[i]
			i++
			//cout << minlen << endl;
		} else if sum < target {
			j++
			if j < len(nums) {
				sum += nums[j]
			}
		} else {
			minlen = min(minlen, j-i+1)
			//cout << minlen << endl;
			sum -= nums[i]
			i++
			j++
			if j < len(nums) {
				sum += nums[j]
			}

		}

	}

	if minlen == math.MaxInt64 {
		return 0
	}
	return minlen
}

func TestLC209(t *testing.T) {
	n := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	minlen := minSubArrayLen(213, n)
	fmt.Println(minlen)
}

func Test8080Read(t *testing.T) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		b := make([]byte, 1024)
		conn.Read(b)
		log.Println(string(b))
	}
}

func Test68__(t *testing.T) {
}

func Test1212(t *testing.T) {
	n := []int{-4, -1, 0, 3, 10}
	absSort(n)
	fmt.Println(n)
}

func absSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var abs = func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	var pow2abs = func(i int) int {
		return int(math.Pow(float64(abs(i)), 2))
	}

	l := 0
	r := len(nums) - 1
	p := nums[0]

	for l < r {
		for l < r && pow2abs(nums[r]) > pow2abs(p) {
			r--
		}
		nums[l] = nums[r]

		for l < r && pow2abs(nums[l]) <= pow2abs(p) {
			l++
		}
		nums[r] = nums[l]
	}

	nums[l] = p
	absSort(nums[:l])
	absSort(nums[l+1:])
}

func Test_lc977(t *testing.T) {
	sortedSquares([]int{-5, -3, -2, -1})
}

func sortedSquares(nums []int) []int {
	var d int // 正负分界点

	for d < len(nums)-1 && nums[d] < 0 {
		d++
	}

	n := make([]int, len(nums), len(nums))

	i := d - 1
	var j int
	for i >= 0 && d < len(nums) {
		if nums[i]*nums[i] < nums[d]*nums[d] {
			n[j] = nums[i] * nums[i]
			i--
			j++
		} else {
			n[j] = nums[d] * nums[d]
			d++
			j++
		}
	}

	// 追加剩余
	for d < len(nums) {
		n[j] = nums[d] * nums[d]
		j++
		d++
	}

	for i >= 0 {
		n[i] = nums[i] * nums[i]
		i--
		j++
	}

	return n
}

func TestMoreDial(t *testing.T) {
	var wg sync.WaitGroup
	c := 100
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", ":8080")
			if err != nil {
				log.Fatalln(err)
			}

			rb := make([]byte, 256)
			n, err := conn.Write([]byte("conn success"))
			if err != nil {
				log.Println("write error")
			}
			fmt.Printf("[%v] write %d bytes\n", conn.LocalAddr(), n)

			n, err = conn.Read(rb)
			if err != nil {
				log.Println("write error")
			}
			//log.Printf("read %d bytes\n", n)

			fmt.Printf("[%v] read content: %v\n", conn.LocalAddr(), string(rb))
		}()
	}
	wg.Wait()
}

func TestStringSub(t *testing.T) {
	s := "1234"
	s1 := s[1:2]
	fmt.Println(s1)
}

func lc43(s1, s2 string) string {
	var (
		flag int    // 保存进位
		res  string //
		pown int    // 确定要补几个 0
	)

	for i := len(s1) - 1; i >= 0; i-- {
		var r strings.Builder
		for j := len(s2) - 1; j >= 0; j-- {
			v1 := int(s1[i] - '0')
			v2 := int(s2[j] - '0')

			tmp := v1*v2 + flag
			flag = tmp / 10
			tmp %= 10

			r.WriteByte(byte(tmp + '0'))
		}

		if flag != 0 {
			r.WriteByte(byte(flag) + '0')
			flag = 0
		}

		rstr := reverseStr(r.String())

		// 		45
		//	*  123
		// --------------	每步后面补 n-1 个 0，n 为 位数
		//     135				135	 补 0 个
		//     90         =>  	900	 补 1 个
		//    45			   4500  补 2 个
		for i := 0; i < pown; i++ {
			rstr += "0"
		}
		pown++

		r1 := strAdd(res, rstr)
		//fmt.Println(r1)
		res = r1
	}

	return res
}

func reverseStr(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func strAdd(s1, s2 string) string {
	var (
		flag int
		res  strings.Builder
	)
	for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var v1, v2 int
		if i >= 0 {
			v1 = int(s1[i] - '0')
		} else {
			v1 = 0
		}

		if j >= 0 {
			v2 = int(s2[j] - '0')
		} else {
			v2 = 0
		}

		sum := v1 + v2 + flag
		flag = sum / 10
		sum %= 10

		res.WriteByte(byte(sum) + '0')
	}

	if flag != 0 {
		res.WriteByte(byte(flag) + '0')
	}

	str := reverseStr(res.String())
	return str
}

func TestLc43(t *testing.T) {
	fmt.Println(lc43("123", "45"))
	fmt.Println(lc43("2", "3"))
	fmt.Println(lc43("123", "20"))
	fmt.Println(lc43("0", "0"))
}

func multiply(num1 string, num2 string) string {
	n1l := len(num1)
	n2l := len(num2)

	n := make([]int, n1l+n2l)

	for i := n1l - 1; i >= 0; i-- {
		v1 := num1[i] - '0'
		for j := n2l - 1; j >= 0; j-- {
			v2 := num2[j] - '0'
			n[i+j+1] += int(v1) * int(v2)
			//fmt.Println(n)
		}
	}

	for i := len(n) - 1; i >= 0; i-- {

	}

	return ""
}

func TestEchoClient(t *testing.T) {
	var wg sync.WaitGroup
	var handler = func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", ":5000")
		if err != nil {
			log.Fatalln(err)
		}

		msg := "message: 123"

		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Println("write error: ", err)
			return
		}
		log.Printf("[%v] send msg: %v\n", conn.LocalAddr(), msg)

		b := make([]byte, 1024)
		_, err = conn.Read(b)
		if err != nil {
			log.Println("read error: ", err)
			return
		}

		log.Printf("[%v] recv msg: %v\n", conn.LocalAddr(), string(b))
	}

	n := 100
	wg.Add(n)

	for i := 0; i < n; i++ {
		go handler()
	}
	wg.Wait()
}

type LFUCache struct {
	cm  map[int]*list.List    // 记录出现次数
	m   map[int]*list.Element // 映射到链表的节点
	cap int                   // 容量
	min int                   // 最少使用次数
}

type node struct {
	key int
	val int
	cnt int // 使用次数
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cm:  make(map[int]*list.List),
		m:   make(map[int]*list.Element),
		cap: capacity,
		min: 1,
	}
}

func (c *LFUCache) Get(key int) int {
	if e, ok := c.m[key]; ok {
		n := e.Value.(*node)

		// 将 n 从旧的链表中移除，因为 get() 后其使用次数增加了
		if c.cm[n.cnt] != nil {
			c.cm[n.cnt].Remove(e)
		}

		// 如果移除后旧链表为空，则更新全局最小使用次数 min
		if c.cm[n.cnt] == nil || c.cm[n.cnt].Len() == 0 {
			// 链表为空了，cm 保存也就没意义了，可以从 cm 中删除
			delete(c.cm, n.cnt)
			if c.min == n.cnt {
				c.min++
			}
		}

		// 更新 n 的使用次数
		n.cnt++

		// 获取新的链表，并将 n 添加到头部
		nl := c.cm[n.cnt]
		if nl == nil {
			nl = list.New()
			c.cm[n.cnt] = nl
		}
		nl.PushFront(n)

		return e.Value.(*node).val
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	// 如果添加该节点后容量溢出，则需要淘汰使用次数最少的节点
	if len(c.m) == c.cap {
		// 根据 min 字段，取出最小使用次数对应的链表
		l := c.cm[c.min]
		if l != nil /*&& l.Len() > 0*/ {
			//fmt.Println(l.Back())
			// 该链表的最后一个即为最长时间未使用的节点，将其移除
			rm := l.Remove(l.Back())
			// 同时也从 m 中移除
			delete(c.m, rm.(*node).key)
		}
	}

	// 如果该 key 已经存在
	if no, ok := c.m[key]; ok {
		n := no.Value.(*node)

		if c.cm[n.cnt] != nil {
			c.cm[n.cnt].Remove(no)
		}

		if c.cm[n.cnt] == nil || c.cm[n.cnt].Len() == 0 {
			delete(c.cm, n.cnt)
			if c.min == n.cnt {
				c.min++
			}
		}

		n.cnt++       // 使用次数 +1
		n.val = value // 更新 val

		ele := c.cm[n.cnt].PushFront(n)
		c.m[key] = ele
		return
	}

	// key 不存在
	n := &node{
		key: key,
		val: value,
		cnt: 1,
	}

	if _, ok := c.cm[1]; !ok {
		c.cm[1] = list.New()
	}

	l := c.cm[1]
	// 添加到计数链表中
	node := l.PushFront(n)
	// 添加到 m 中
	c.m[key] = node
	// 更新全局最小使用次数 min 为 1
	c.min = 1
}

// 将节点 n 从 c.cm[oldCnt] 所指的链表中删除，并添加到 c.cm[newCnt] 所指
// 的链表中，该项操作适用于某一节点的使用次数增加时
func (c *LFUCache) moveToAnotherList(oldCnt, newCnt int, n *list.Element) {
	if _, ok := c.cm[oldCnt]; !ok {
		c.cm[oldCnt] = list.New()
	}

	if _, ok := c.cm[newCnt]; !ok {
		c.cm[newCnt] = list.New()
	}

	// 从旧链表中删除
	ol := c.cm[oldCnt]
	ol.Remove(n)

	// 添加到新链表的头部
	nl := c.cm[newCnt]
	nl.PushFront(n)
	// 同时更新 n 的使用次数 cnt
	n.Value.(*node).cnt = newCnt
}

func TestLFU(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	fmt.Println(l.Get(1))
	l.Put(3, 3)
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
}

func TestMmap(t *testing.T) {
	//syscall.Mmap(1, 0, 0, syscall.PROT_READ)
	//syscall.Munmap( )
	//syscall.SetNonblock()
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s == "" || goal == "" {
		return true
	}

	var i, j int
	var count int
	for i < len(s) && s[i] != goal[j] {
		i++
	}
	fmt.Println(i)

	for {
		if i == len(s) {
			i = 0
		}
		if j == len(goal) {
			j = 0
		}

		fmt.Println(string(s[i]), string(goal[j]))
		if s[i] != goal[j] {
			return false
		}
		i++
		j++
		count++

		if count == len(s) {
			break
		}
	}

	return true
}

func TestLc796(t *testing.T) {
	//rotateString("abcde", "abced")
	//rotateString("abcde", "cdeab")

	//"bbbacddceeb"
	//"ceebbbbacdd"
	r := rotateString("bbbacddceeb", "ceebbbbacdd")
	fmt.Println(r)

	s := "13"
	s = strings.TrimSpace(s)

}

func TestStringsMap(t *testing.T) {
	s := "a123b"
	s2 := strings.Map(func(r rune) rune {
		if r == 'a' {
			return 'A'
		}
		return r
	}, s)

	fmt.Println(s2)
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	var maxlen int
	var curval int
	var res int

	for _, v := range nums {
		m[v] = true
	}

	for k := range m {
		if _, ok := m[k-1]; ok {
			continue
		} else {
			maxlen++
			curval = k

			for m[curval+1] {
				maxlen++
				curval++
			}
		}
		res = max(res, maxlen)
	}

	return res
}

func Test_iajsod(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(addr.String())
}

func TestPanic_Recover(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			debug.PrintStack()
			fmt.Println("somewhere panic")
		}
	}()

	var f = func() {
		panic("panic!")
	}

	f()
}

func movingCount(m int, n int, k int) int {
	visit := make([][]bool, m)
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, n)
	}

	return dfs(m, n, k, 0, 0, visit)
}

func sum(m, n int) int {
	return m/10 + m%10 + n/10 + n%10
}

func dfs(m, n, k, i, j int, visit [][]bool) int {
	if i > m || j > n || visit[i][j] || sum(i, j) > k {
		return 0
	}
	visit[i][i] = true
	return 1 + dfs(m, n, k, i+1, j, visit) + dfs(m, n, k, i-1, j, visit) + dfs(m, n, k, i, j+1, visit) + dfs(m, n, k, i, j-1, visit)
}

func exist(board [][]byte, word string) bool {
	var res bool
	backtrack111(board, word, "", 0, 0, &res)

	return res
}

func backtrack111(board [][]byte, word, cur string, i, j int, res *bool) {
	if len(cur) == len(word) {
		fmt.Println(cur)
		if cur == word {
			*res = true
			return
		}
		return
	}

	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[0]) {
		return
	}

	cur += (string)(board[i][j])

	backtrack111(board, word, cur, i+1, j, res)
	backtrack111(board, word, cur, i-1, j, res)
	backtrack111(board, word, cur, i, j-1, res)
	backtrack111(board, word, cur, i, j+1, res)

	cur = cur[:len(cur)-1]

}

func TestOffer12(t *testing.T) {
	b := [][]byte{
		{'a', 'b'},
	}
	exist(b, "ba")
}

func TestUnsafeSizeOf(t *testing.T) {
	var i uint32
	sizeof := unsafe.Sizeof(&i)
	fmt.Println(sizeof)

	s := []int{1}
	fmt.Println(s[1:])

}

func TestScan(t *testing.T) {

}

// errors.New() 是否输出堆栈信息？
func TestErrorsIsPrintStack(t *testing.T) {
	i := 1000

	f := func(a int) error {
		if a < 100 {
			return errors.New("error!")
		}
		return nil
	}

	if err := f(i); err != nil {
		fmt.Printf("%+v", err)
		debug.PrintStack()
		return
	}
}

func TestHttp2(t *testing.T) {
	bufio.NewWriter(os.Stdin)
}

func TestReadAt(t *testing.T) {

}

func TestList(t *testing.T) {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(l.Front())
		}()
		go func(j int) {
			l.PushFront(j)
			fmt.Println("push ", j)
		}(i)
	}
}

func TestAtomicUintSub(t *testing.T) {
	var i uint64 = 100
	var e int64 = -10

	atomic.AddUint64(&i, ^uint64(-e-1))
	fmt.Println(i)
}

func TestSearch(t *testing.T) {
	sort.SearchInts(nil, 1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	s := &ListNode{Next: head}

	var (
		pre   = s
		next  = pre
		start *ListNode
		end   = pre
	)

	for pre != nil {
		start = pre.Next
		for i := 0; i < k; i++ {
			if end == nil {
				break
			}
			end = end.Next
		}
		next = end.Next
		end.Next = nil // 断链

		rh, rt := reverse720(start) // 反转 start - end
		pre.Next = rh
		rt.Next = next

		pre = rt
		end = rt

	}

	return s.Next
}

func reverse720(head *ListNode) (rhead, rtail *ListNode) {
	var (
		cur        = head
		prev, next *ListNode
	)

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev, head
}

func printList(head *ListNode) {
	h := head
	for h != nil {
		fmt.Printf("%v ->", head.Val)
		h = h.Next
	}
	fmt.Println("")
}

func Test720(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}

	reverseKGroup(l, 2)
}

type myHeap []int

func (h myHeap) Len() int {
	return len(h)
}

func (h myHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h myHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {
	p := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return p
}

func (h myHeap) Front() int {
	return h[len(h)-1]
}

func TopK(arr []int, k int) []int {
	h := new(myHeap)

	for i := 0; i < len(arr); i++ {
		if h.Len() < k {
			h.Push(arr[i])
		} else {
			if h.Front() > arr[i] {
				heap.Pop(h)
				heap.Push(h, arr[i])
			}
		}
	}

	return *h
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	q := arr[0]
	l, r := 0, len(arr)-1

	for l < r {
		for r > l && arr[r] > q {
			r--
		}
		arr[l] = arr[r]

		for r > l && arr[l] <= q {
			l++
		}
		arr[r] = arr[l]
	}

	arr[l] = q

	quickSort(arr[:r])
	quickSort(arr[r+1:])
}

func TestQuickSort(t *testing.T) {
	n := []int{2, 5, 3, 5, 1, -1, 9}
	quickSort(n)
	fmt.Println(n)
	// http.ReadRequest()
}

var i64 int64

func TestCas(t *testing.T) {
	for i1 := 0; i1 < 20; i1++ {
		go fn(i1)
	}

	time.Sleep(time.Second * 30)
}

func fn(seq int) {
	fmt.Printf("goroutine[%d] \n", seq)
	// 自旋 3 次
	for i1 := 0; i1 < 3; i1++ {
		if atomic.CompareAndSwapInt64(&i64, i64, i64+1) {
			fmt.Printf("goroutine[%d] cas ok, i: %d \n", seq, i64)
			break
		}
	}
	fmt.Printf("goroutine[%d] done. \n", seq)
}

func fmt729() {
	// 标准格式
	for i := 0 + 1; i < 50-1; i++ {

	}
}

func TestListen7070(t *testing.T) {
	listen, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(conn.RemoteAddr())
		conn.Write([]byte("ok"))
	}
}

func TestDial7070(t *testing.T) {
	conn, err := net.Dial("tcp", ":7070")
	if err != nil {
		log.Fatalln(err)
	}

	b := make([]byte, 1024)
	conn.Read(b)
	//br := bufio.NewReader(conn)
	//val, _, _ := br.ReadLine()
	fmt.Println(string(b))
}

func reverseStr1(s string, k int) string {
	var (
		lo, hi int
		res    string
	)
	for {
		for i := 0; i < (2*k)-1; i++ {
			if hi < len(s) {
				hi++
			}
		}
		strlen := hi - lo + 1

		if strlen < k {
			res += reverse1(s, lo, hi)
		} else if strlen < 2*k && strlen >= k {
			res += reverse1(s, lo, lo+1)
			return res // 不足 2k 个，说明到最后一部分了
		} else if strlen >= 2*k {
			res += reverse1(s, lo, lo+1)
		}
		lo = hi + 1
		fmt.Println(res)
	}

	return res
}

func reverse1(s string, i, j int) string {
	s1 := []byte(s)
	for j < len(s) && i < j {
		s1[i], s1[j] = s1[j], s1[i]
		i++
		j--
	}

	return string(s1[i : j+1])
}

func TestStrReverse1(t *testing.T) {
	reverseStr1("abcdefg", 2)
}

func format() {
	// p = 1
	pr := findPath(root, p, &path, &res, &flag)
	for i := 0; i < len(pr); i++ {
		fmt.Printf("%v ", pr[i]) // Output: [3, 1]
	}
	fmt.Println()

	flag = false
	// p = 4
	qr := findPath(root, q, &path, &res, &flag)

	for i := 0; i < len(qr); i++ {
		fmt.Printf("%v ", qr[i]) // Output: [3,5,2,4]
	}
}

func findPath(root, need *TreeNode, path, res *[]*TreeNode, flag *bool) {
	if root == nil || *flag {
		return
	}
	*path = append(*path, root)
	if root == need {
		*flag = true
		// copy 一个新切片，防止多个 res 指向同一个 path
		news := make([]*TreeNode, len(*path))
		copy(news, *path)
		*res = news
		return
	}

	findPath(root.Left, need, path, res, flag)
	findPath(root.Right, need, path, res, flag)

	*path = (*path)[:len(*path)-1]
	return
}

func main() {
	var path, pp, qq []*TreeNode
	var flag bool
	findPath(root, p, &path, &pp, &flag)
	// 重置参数
	flag = false
	path = path[0:0]
	findPath(root, q, &path, &qq, &flag)
	// 此时 pp 和 qq 已经保存了结果
}
