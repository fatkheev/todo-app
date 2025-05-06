package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Связь пользователей и списков (у одного пользователя несколько списков)
type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Связь списка и предметов (у одного списка несколько предметов)
type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
