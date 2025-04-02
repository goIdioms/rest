package models

type TodoItem struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
	Done        bool   `json:"done"`
}

type TodoList struct {
	Id          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Discription string `json:"discription"`
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
