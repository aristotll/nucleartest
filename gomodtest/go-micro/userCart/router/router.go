package router

import (
	"github.com/gin-gonic/gin"
	"go-micro/userCart/handler"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/user/cart", handler.GetCartInfo)
	r.POST("/user/carts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "123",
		})
	})
	r.POST("/user/cart/json", handler.GetCartInfoById)
	return r
}