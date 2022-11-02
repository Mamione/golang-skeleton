package main

import (
	"nameless/api/account"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", account.CreateAccount)
	r.Run() // listen and serve on 0.0.0.0:8080
}