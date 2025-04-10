package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	listsItemTable = "lists_item"
	userListsTable = "user_lists"
	todoListsTable = "todo_lists"
	usersTable     = "users"
	todoItemsTable = "todo_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	db.Ping()
	return db, nil
}
