package todo

import "context"

type TodoRepositiry interface {
	GetTodo(ctx context.Context, id int64) (Todo, error)
}
