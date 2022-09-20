package main

type S[T any] struct{ x T }

type T[T any] struct {
	// X T // 这里的 T 是 struct T 而不是泛型 T any
}

//func (T[T]) m() {} // error: T is not a generic type
// 看网上的更新日志（非官方英文）说 go1.19 不会报错了，但是这样的代码有什么用吗？泛型都没有用到
// 而且跑了一下，怎么 1.19 依然会报错。。。

type MAP[K comparable, V any] map[K]V

func newMap[K comparable, V any](cap int64) MAP[K, V] {
	return make(MAP[K, V], cap)
}

func (m MAP[K, V]) Insert(k K, v V) {
	m[k] = v
}

func (m MAP[K, V]) Find(k K) (v V, ok bool) {
	v, ok = m[k]
	return
}

func (m MAP[K, V]) Delete(k K) {
	delete(m, k)
}

func main() {
	var t S[int64]
	t.x = 123
	//t.x = ""
	println(t.x)

	var m = make(MAP[string, string])
	m["1"] = "2"

	m1 := newMap[int64, int64](10)
	m1.Insert(1, 2)
	println(m1.Find(1))
	m1.Delete(1)
	println(m1.Find(1))
}
