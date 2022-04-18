package infra

import (
	"context"
	"github.com/Hiroya3/clean-architecture-practice/grpc/internal/domain/todo"
)

type TodoRepository struct {
	todoGateway TodoGateway
}

func (t *TodoRepository) GetTodo(ctx context.Context, id int64) (todo.Todo, error) {
	return t.todoGateway.Select(ctx, id)
}
