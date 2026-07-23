package service

import (
	"time"
	"sync"

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

func (s *TaskService) GetTasks(query dto.TaskQuery) ([]model.Task, int64, error) {

	var (
		tasks []model.Task
		total int64

		taskErr error
		totalErr error
	)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {

		defer wg.Done()

		tasks, taskErr = s.repo.GetTasks(query)

	}()

	go func() {

		defer wg.Done()

		total, totalErr = s.repo.CountTasks(query)

	}()

	wg.Wait()

	if taskErr != nil {
		return nil, 0, taskErr
	}

	if totalErr != nil {
		return nil, 0, totalErr
	}

	return tasks, total, nil
}