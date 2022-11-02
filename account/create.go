package account

import "github.com/gin-gonic/gin"

func create(username string, password string) bool {
	return true
}

func CreateAccount(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello!",
	})
}