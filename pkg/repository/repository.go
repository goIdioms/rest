package repository

import (
	"github.com/goIdioms/rest/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string) (models.User, error)
}

type TodoList interface {
	Create(userID int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, id int) (models.TodoList, error)
	Delete(userId, id int) error
	Update(userId, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
