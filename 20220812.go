package main

type slices[T any] []T

func fn() interface{} {
    return 1
}

func fn_[T any](s slices[T]) {
    _ = s    
}

func main() {
    fn()
    var s []int
    fn_(slices[int](s))
    _ = s
}
