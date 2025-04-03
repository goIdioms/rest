package repository

import (
	"fmt"
	"strings"

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

func (r *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.discription FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, userListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (models.TodoList, error) {
	var list models.TodoList

	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.discription FROM %s tl
								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		todoListsTable, userListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodoListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, userListsTable)
	_, err := r.db.Exec(query, userId, id)
	return err
}

func (r *TodoListPostgres) Update(userId, listId int, input models.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Discription != nil {
		setValues = append(setValues, fmt.Sprintf("discription=$%d", argId))
		args = append(args, *input.Discription)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, userListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
