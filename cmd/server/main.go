package main

import (
	"todo-api/config"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Todo API Running",
		})
	})

	r.Run(":8080")
}