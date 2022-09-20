package main

import (
    "golang.org/x/exp/slices"
    "fmt"
)

func main() {
    s := []int{1, 2, 3, 4, 5, 6}
    s = slices.Insert(s, 2, 7, 8, 9)
    fmt.Printf("s=%v, after insert(index=2, val=7,8,9), s=%v\n", s, s)

    s = slices.Delete(s, 2, 5)
    fmt.Printf("s=%v, after delete(2, 5), s=%v\n",s, s)

    if !slices.IsSorted(s) {
        fmt.Printf("s=%v is not sorted\n", s)
        slices.Sort(s)
        fmt.Printf("after sort, s=%v\n", s)
    }

    si := 2
    idx, find := slices.BinarySearch(s, si)
    if find {
        fmt.Printf("find %v in s=%v, index=%v\n", si, s, idx)
    } else {
        fmt.Printf("not found %v\n", si)
    }
}
