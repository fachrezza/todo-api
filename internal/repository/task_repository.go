package repository

import (
	"strings"


	"github.com/fachrezza/todo-api/config"
	"github.com/fachrezza/todo-api/internal/model"
	"github.com/fachrezza/todo-api/internal/dto"
)

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) Create(task *model.Task) error {
	return config.DB.Create(task).Error
}

func (r *TaskRepository) GetTasks(query dto.TaskQuery) ([]model.Task, error) {

	var tasks []model.Task

	db := config.DB.Model(&model.Task{})

	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	if query.Search != "" {

		search := "%" + strings.ToLower(query.Search) + "%"

		db = db.Where(
			"LOWER(title) LIKE ? OR LOWER(description) LIKE ?",
			search,
			search,
		)
	}

	offset := (query.Page - 1) * query.Limit

	err := db.
		Limit(query.Limit).
		Offset(offset).
		Order("created_at desc").
		Find(&tasks).Error

	return tasks, err
}

func (r *TaskRepository) CountTasks(query dto.TaskQuery) (int64, error) {

	var total int64

	db := config.DB.Model(&model.Task{})

	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}

	if query.Search != "" {

		search := "%" + strings.ToLower(query.Search) + "%"

		db = db.Where(
			"LOWER(title) LIKE ? OR LOWER(description) LIKE ?",
			search,
			search,
		)
	}

	err := db.Count(&total).Error

	return total, err
}

func (r *TaskRepository) GetByID(id string) (*model.Task, error) {

	var task model.Task

	err := config.DB.First(&task, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Update(task *model.Task) error {

	return config.DB.Save(task).Error

}

func (r *TaskRepository) Delete(task *model.Task) error {

	return config.DB.Delete(task).Error

}