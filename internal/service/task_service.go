package service

import (
	"time"

	"github.com/fachrezza/todo-api/internal/dto"
	"github.com/fachrezza/todo-api/internal/model"
	"github.com/fachrezza/todo-api/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) Create(req dto.CreateTaskRequest) (*model.Task, error) {

	dueDate, err := time.Parse("2006-01-02", req.DueDate)

	if err != nil {
		return nil, err
	}

	task := model.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     dueDate,
	}

	err = s.repo.Create(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}