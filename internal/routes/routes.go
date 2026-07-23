package routes

import (
	"github.com/fachrezza/todo-api/internal/handler"
	"github.com/fachrezza/todo-api/internal/repository"
	"github.com/fachrezza/todo-api/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	repo := repository.NewTaskRepository()

	svc := service.NewTaskService(repo)

	taskHandler := handler.NewTaskHandler(svc)

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
}