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

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, id int) (models.TodoList, error) {
	return s.repo.GetById(userId, id)
}

func (s *TodoListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *TodoListService) Update(userId, listId int, input models.UpdateListInput) error {
	return s.repo.Update(userId, listId, input)
}
