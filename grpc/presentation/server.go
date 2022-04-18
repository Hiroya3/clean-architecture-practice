package presentation

import (
	"context"
	"github.com/Hiroya3/clean-architecture-practice/grpc/gen"
	"github.com/Hiroya3/clean-architecture-practice/grpc/internal/service"
)

type GprcServer struct {
	TodoService service.TodoService
}

func NewGrpcServer(TodoService service.TodoService) *GprcServer {
	return &GprcServer{
		TodoService: TodoService,
	}
}

func (s *GprcServer) GetTodoV1(ctx context.Context, request *gen.GetTodoV1RequestPayload) (*gen.GetTodoV1ResponsePayload, error) {
	todoID := request.GetId()

	got, err := s.TodoService.Get(ctx, todoID)
	if err != nil {
		return nil, err
	}

	response := &gen.GetTodoV1ResponsePayload{
		TodoName: got.Text,
	}

	return response, nil
}
