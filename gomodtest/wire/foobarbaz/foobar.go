package foobarbaz

import (
	"errors"
	"github.com/google/wire"
	"strconv"
	"strings"
)

// 提供者（provider）集合
var SuperSet = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz)

// 也可以添加一个集合到另一个集合
var MegaSet = wire.NewSet(SuperSet)
var BazzNeed = wire.NewSet(ProvideFoo, ProvideBar, ProvideBaz, ProvideBazz)

// demo 中的依赖关系：Baz -> Bar -> Foo

type Foo struct {
	X int
}

// ProvideFoo 是 Foo 的提供者，它没有依赖
func ProvideFoo(x int) *Foo {
	return &Foo{X: x}
}

type Bar struct {
	X int
}

// ProvideBar 是 Bar 的提供者，它依赖于 Foo
func ProvideBar(foo *Foo) *Bar {
	return &Bar{X: foo.X}
}

type Baz struct {
	X int
}

// ProvideBax 是 Baz 的提供者，它依赖于 Bar，而 Bar 又依赖于 Foo，所以 Baz 也依赖于 Foo
// 提供者也可以返回 error
func ProvideBaz(bar *Bar) (*Baz, error) {
	if bar.X == 0 {
		return &Baz{}, errors.New("bar.X == 0")
	}
	return &Baz{X: bar.X}, nil
}

// 一个组合了多个结构体的结构体
type Bazz struct {
	Baz_ *Baz
	Bar_ *Bar
	Foo_ *Foo
}

func ProvideBazz(baz *Baz, bar *Bar, foo *Foo) (*Bazz, error) {
	bzz := &Bazz{
		Baz_: baz,
		Bar_: bar,
		Foo_: foo,
	}
	return bzz, nil
}

func (b *Bazz) String() string {
	var sb strings.Builder
	sb.WriteString("foo.x = " + strconv.Itoa(b.Foo_.X) + "\n")
	sb.WriteString("bar.x = " + strconv.Itoa(b.Bar_.X) + "\n")
	sb.WriteString("baz.x = " + strconv.Itoa(b.Baz_.X) + "\n")
	return sb.String()
}
