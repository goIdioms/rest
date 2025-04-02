package repository

import (
	"fmt"

	"github.com/goIdioms/rest/pkg/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list models.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.Fatal(err)
	}

	var id int
	listQuery := fmt.Sprintf("INSERT INTO %s (title, discription) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(listQuery, list.Title, list.Discription)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		logrus.Fatal(err)
	}

	usersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", userListsTable)
	_, err = tx.Exec(usersListQuery, userID, id)
	if err != nil {
		tx.Rollback()
		logrus.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		logrus.Fatal(err)
	}

	return id, err
}
