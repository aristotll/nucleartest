package list

type list[T any] struct {
	root *Element[T]
	len  int64
}

func New[T any]() *list[T] {
	return new(list[T]).Init()
}

func (l *list[T]) Len() int64 {
	return l.len
}

func (l *list[T]) Init() *list[T] {
	l.root = &Element[T]{}
	l.root.next = l.root
	l.root.prev = l.root
	l.len = 0
	return l
}

func (l *list[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *list[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *list[T]) insert(e, at *Element[T]) *Element[T] {
	e.next = at.next
	at.next.prev = e
	at.next = e
	e.prev = at
	l.len++
	e.list = l
	return e
}

func (l *list[T]) insertValue(val T, at *Element[T]) T {
	return l.insert(&Element[T]{Value: val}, at).Value
}

func (l *list[T]) remove(at *Element[T]) *Element[T] {
	at.prev.next = at.next
	at.next.prev = at.prev
	at.next = nil
	at.prev = nil
	l.len--
	at.list = nil
	return at
}

func (l *list[T]) PushBack(v T) T {
	return l.insertValue(v, l.root.prev)
}

func (l *list[T]) PushFront(v T) T {
	return l.insertValue(v, l.root)
}

func (l *list[T]) Remove(e *Element[T]) T {
	if l.len == 0 {
		panic("can not remove element with empty list")
	}
	if e == nil {
		panic("e is nil")
	}
	if e.list != l {
		panic("e does not belong this list")
	}
	return l.remove(e).Value
}

func (l *list[T]) Range(fn func(T)) {
	for i := l.Front(); i != nil; i = i.Next() {
		fn(i.Value)
	}
}

func (l *list[T]) PopFront() T {
	return l.Remove(l.Front())
}

func (l *list[T]) PopBack() T {
	return l.Remove(l.Back())
}

type Element[T any] struct {
	next  *Element[T]
	prev  *Element[T]
	list  *list[T]
	Value T
}

func (e *Element[T]) Next() *Element[T] {
	if v := e.next; e.list != nil && v != e.list.root {
		return v
	}
	return nil
}

func (e *Element[T]) Prev() *Element[T] {
	if v := e.prev; e.list != nil && v != e.list.root {
		return v
	}
	return nil
}
