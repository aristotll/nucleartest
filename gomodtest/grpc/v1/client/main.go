package main

import (
	"gRPC/v1/client/clientFunc"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//clientFunc.GetProdInfo()
	//clientFunc.GetStock()
	//clientFunc.GetStocks()
	//
	//clientFunc.CreateOrder()
	//
	//clientFunc.GetUserScores()
	//clientFunc.GetUserScoresByServerStream()
	//clientFunc.GetUserScoresByClientStream()
	// clientFunc.GetUserScoresByTwoWayStream()

	r := gin.Default()
	r.POST("/spike", clientFunc.SubmitOrder)

	err := r.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}
