package infra

import (
	"context"
	"github.com/Hiroya3/clean-architecture-practice/grpc/internal/domain/todo"
)

type TodoGateway interface {
	Select(ctx context.Context, id int64) (todo.Todo, error)
}

type Gateway struct {
}

func (g *Gateway) Select(ctx context.Context, id int64) (todo.Todo, error) {
	return todo.Todo{
		ID:     id,
		Text:   "Text",
		Done:   false,
		UserId: 1,
	}, nil
}
