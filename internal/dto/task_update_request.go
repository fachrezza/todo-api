package dto

type UpdateTaskRequest struct {
	Title       string `json:"title" binding:"required,max=255"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required,oneof=pending completed"`
	DueDate     string `json:"due_date" binding:"required"`
}