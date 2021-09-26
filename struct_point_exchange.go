package main

import (
    "fmt"
)

type T_T struct {
    name string
    age int
}

func exchange(t1, t2 *T_T) {
    *t1, *t2 = *t2, *t1
}

func exNo(t1, t2 *T_T) {
    t1, t2 = t2, t1
}

// 通过指针接收者交换

type tt struct {
    a, b *T_T
}

func (t *tt) exchange() {
    t.a, t.b = t.b, t.a
}

func (t *tt) exc() {
    a := t.a
    b := t.b
    *a, *b = *b, *a
}

// 面向过程 test
func test1() {
    t1 := T_T{"zhang", 123}
    t2 := T_T{"wang", 555}
    fmt.Printf("before: %v %v\n", t1, t2)
    //exchange(&t1, &t2)
    exNo(&t1, &t2)
    fmt.Printf("after: %v %v\n", t1, t2)
}

func test2() {
    t := &tt{
        a: &T_T{"123", 18},
        b: &T_T{"456", 22},
    }
    fmt.Printf("before: %+v %+v\n", t.a, t.b)
    

    t.exchange()
    //t.exc()
    fmt.Printf("after: %+v %+v\n", t.a, t.b)
}

func main() {
    test2()
}
