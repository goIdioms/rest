package service

import (
	"github.com/goIdioms/rest/pkg/models"
	"github.com/goIdioms/rest/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list models.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}
