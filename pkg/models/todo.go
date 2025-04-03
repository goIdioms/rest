package models

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
	Discription *string `json: "discription"`
}
