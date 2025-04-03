package models

import "errors"

type TodoItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Done        bool   `json:"done"`
}

type TodoList struct {
	Id          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Discription string `json:"discription" db:"discription"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type ListsItem struct {
	Id     int
	ItemId int
	ListId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
