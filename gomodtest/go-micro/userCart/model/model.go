package model

type UserCart struct {
	Id int
	ProductName string
	// 数量
	Quantity int
	// 总价格
	TotalCost float64
}
