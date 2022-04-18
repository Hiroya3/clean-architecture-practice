package service

import (
	"context"
	"github.com/Hiroya3/clean-architecture-practice/grpc/internal/domain/todo"

	"go.uber.org/zap"
)

type TodoService struct {
	logger         zap.Logger
	todoRepository todo.TodoRepositiry
}

func (t *TodoService) Get(ctx context.Context, id int64) (todo.Todo, error) {

	got, err := t.todoRepository.GetTodo(ctx, id)
	if err != nil {
		t.logger.Warn("fail to get todo", zap.Error(err))
		return todo.Todo{}, err
	}

	return got, nil
}
