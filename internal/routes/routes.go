package routes

import (
    "github.com/gin-gonic/gin"

    "github.com/fachrezza/todo-api/config"
    "github.com/fachrezza/todo-api/internal/handler"
    "github.com/fachrezza/todo-api/internal/middleware"
    "github.com/fachrezza/todo-api/internal/repository"
    "github.com/fachrezza/todo-api/internal/service"
)

func SetupRoutes(r *gin.Engine) {

    taskRepo := repository.NewTaskRepository()
    taskService := service.NewTaskService(taskRepo)
    taskHandler := handler.NewTaskHandler(taskService)

    userRepo := repository.NewUserRepository(config.DB)
    authService := service.NewAuthService(userRepo)
    authHandler := handler.NewAuthHandler(authService)

    r.POST("/register", authHandler.Register)
    r.POST("/login", authHandler.Login)

    task := r.Group("/tasks")
    task.Use(middleware.JWTMiddleware())

    task.POST("", taskHandler.CreateTask)
    task.GET("", taskHandler.GetTasks)
    task.GET("/:id", taskHandler.GetTaskByID)
    task.PUT("/:id", taskHandler.UpdateTask)
    task.DELETE("/:id", taskHandler.DeleteTask)
}