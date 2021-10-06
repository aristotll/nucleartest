package main

func main() {
	m := make(map[int]int)
	m[1] = 2
	//fmt.Printf("m: %v\n", m)

	m2 := make(map[int]int)
	m2[1] = m[1]
}
