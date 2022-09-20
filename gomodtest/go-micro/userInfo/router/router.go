package router

import (
	"github.com/gin-gonic/gin"
	"go-micro/userInfo/handler"
	"net/http"
)

func Router() *gin.Engine{
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "user  success",
		})
	})
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg": "book  success",
		})
	})
	r.GET("/service/random", handler.GetServiceByRandom)
	r.GET("/service/round", handler.GetServiceByPoll)
	r.GET("/call/basic", handler.BasicCall)
	r.GET("/call/plugins", handler.CallByPlugins)
	r.GET("/call/grpc", handler.CallBygRPC)

	return r
}
