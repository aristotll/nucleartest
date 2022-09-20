package main

import "fmt"

type UserCart struct {
	Id int
	ProductName string
	// 数量
	Quantity int
	// 总价格
	TotalCost float64
}

func newT() {
	u := new([]UserCart)
	// 通过 new 创建的 切片结构体，如何接收值
	u = &[]UserCart{
		{
			ProductName: "iphone1",
		},
		{
			ProductName: "iphone2",
		},
	}
	fmt.Println(u)

}

func main() {
	newT()
	//userC := make([]*UserCart, 5)
	//for i := 0; i < 5; i++ {
	//	userC[i] = &UserCart {
	//		Id:          i + 1,
	//		ProductName: "iphone" + string(i),
	//		Quantity:    i * 10,
	//		TotalCost:   float64(i) * 2000,
	//	}
	//}
	//for _, v := range userC{
	//	fmt.Printf("%+v \n", *v)
	//}
	//fmt.Println(userC)
}
