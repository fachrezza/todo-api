package main

import (
	"github.com/fachrezza/todo-api/config"
	"github.com/fachrezza/todo-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}