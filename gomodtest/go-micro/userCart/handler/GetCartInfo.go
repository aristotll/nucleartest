package handler

import (
	"github.com/gin-gonic/gin"
	"go-micro/userCart/model"
	"log"
	"net/http"
	"strconv"
)

// 模拟获取用户购物车信息
func GetCartInfo(c *gin.Context) {
	userC := make([]*model.UserCart, 5)
	for i := 0; i < 5; i++ {
		userC[i] = &model.UserCart {
			Id:          i + 1,
			ProductName: "iphone" + strconv.Itoa(i),
			Quantity:    i * 10,
			TotalCost:   float64(i) * 2000,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "success",
		"data": userC,
	})
}

func GetCartInfoById(c *gin.Context) {
	uc := new(model.UserCart)
	err := c.ShouldBindJSON(&uc)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"Id": uc.Id,
		"productName": uc.ProductName,
	})
}
