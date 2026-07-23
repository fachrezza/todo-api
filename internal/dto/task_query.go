package dto

type TaskQuery struct {
	Page   int
	Limit  int
	Status string
	Search string
}