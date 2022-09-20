package wirefile

import (
	"github.com/google/wire"
	"wire/foobarbaz"
)


// Build() 创建一个 Baz 对象，其参数是 Baz 依赖的所有提供者
func InitializeBaz(x int) (*foobarbaz.Baz, error) {
	wire.Build(foobarbaz.MegaSet) // 可以传入一个提供者集合
	// 传入所需的依赖提供者
	//wirefile.Build(foobarbaz.ProvideBar, foobarbaz.ProvideFoo, foobarbaz.ProvideBaz)
	return &foobarbaz.Baz{}, nil
}

func InitializeBazz(x int) (*foobarbaz.Bazz, error) {
	//wire.Build(foobarbaz.BazzNeed)
	// 可以直接使用下面的简便方法 wire.Struct()，这样就不用编写 ProvideBazz() 了
	wire.Build(
		foobarbaz.ProvideBar,
		foobarbaz.ProvideFoo,
		foobarbaz.ProvideBaz,
		// * 代表自动注入所有属性
		wire.Struct(new(foobarbaz.Bazz), "*"))
	return &foobarbaz.Bazz{}, nil
}
