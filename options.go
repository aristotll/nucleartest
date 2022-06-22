package main

type Option[T any]  chan T

func (o Option[T]) Some(v T) T {
	if o == nil {
		o = make(chan T, 1)
	}
	o <- v
	return v
}

func (o Option[T]) None() bool {
	return len(o) == 0
}

type st struct{}

func (_ *st) Do() {}

func main() {
	var o Option[*st]
	select o {
	case v := <- o:
	}
}
