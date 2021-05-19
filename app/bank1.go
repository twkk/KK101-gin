package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance = 1000

func main() {
	router := gin.Default()
	router.GET("/balance/", getBalance)

	router.Run(":80")
}

func getBalance(context *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
		"saddr":string:string{"hsnId" :"C" ,"addr" :"中正路12號" } ,
	})
}