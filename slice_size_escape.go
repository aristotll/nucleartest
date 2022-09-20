package main

func fn() {
	b := make([]byte, 1000)
	b1 := make([]byte, 1024)

	b[0] = 1
	b1[0] = 1
}

func main() {
	fn()
}
