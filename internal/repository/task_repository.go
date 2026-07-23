package repository

import (
	"github.com/fachrezza/todo-api/config"
	"github.com/fachrezza/todo-api/internal/model"
)

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) Create(task *model.Task) error {
	return config.DB.Create(task).Error
}