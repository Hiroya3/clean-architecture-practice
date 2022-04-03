package todo

import "Hiroya3/clean-architecture-practice/app/domain/user"

type Todo struct {
	ID   ID
	Text Text
	Done Done
	User *user.User
}

type NewTodo struct {
	Text   Text
	UserID user.ID
}
