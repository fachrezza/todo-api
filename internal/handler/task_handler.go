package handler

import (
	"net/http"
	"math"
	"strconv"

	"github.com/fachrezza/todo-api/internal/dto"
	"github.com/fachrezza/todo-api/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {

	var req dto.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	task, err := h.service.Create(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

func (h *TaskHandler) GetTasks(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	query := dto.TaskQuery{

		Page: page,

		Limit: limit,

		Status: c.Query("status"),

		Search: c.Query("search"),
	}

	tasks, total, err := h.service.GetTasks(query)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(http.StatusOK, gin.H{

		"tasks": tasks,

		"pagination": gin.H{

			"current_page": page,

			"total_pages": totalPages,

			"total_tasks": total,
		},
	})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {

	id := c.Param("id")

	task, err := h.service.GetByID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})

		return
	}

	c.JSON(http.StatusOK, task)

}

func (h *TaskHandler) UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var req dto.UpdateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	task, err := h.service.Update(id, req)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"task": task,
	})

}

func (h *TaskHandler) DeleteTask(c *gin.Context) {

	id := c.Param("id")

	err := h.service.Delete(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "Task not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})

}