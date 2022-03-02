package list

import (
	"testing"
)

func TestList_PushBack(t *testing.T) {
	l := New[int64]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	t.Logf("list size: %d \n", l.Len())
	l.Range(func(v int64) {
		t.Logf("%v\n", v)
	})
}

func TestList_Remove(t *testing.T) {
	type s struct {
		a, b int64
	}

	l := New[*s]()
	l.PushBack(&s{1, 2})
	l.PushBack(&s{2, 3})

	l.Range(func(v *s) { t.Logf("%v -> ", v) })
	t.Logf("\n")

	t.Log("pop back")
	l.Remove(l.Back())
	l.Range(func(v *s) { t.Logf("%v -> ", v) })
	t.Logf("\n")

	l.Remove(l.Back())
	l.Range(func(v *s) { t.Logf("%v -> ", v) })
	t.Logf("\n")

	// 测试对空链表调用 remove
	l.Init()
	l.Remove(l.Back())
}

func TestPop(t *testing.T) {
	l := New[int64]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.PopBack()

}
