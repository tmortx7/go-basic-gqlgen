package usecase

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/repository"
)

type todo struct {
	todoRepository repository.Todo
}

// Todo of todocase
type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TodoWhereInput) (*model.TodoConnection, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
	Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error)
}

// NewTodoUsecase returns todo usecse
func NewTodoUsecase(r repository.Todo) Todo {
	return &todo{todoRepository: r}
}

func (u *todo) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	return u.todoRepository.Get(ctx, id)
}

func (u *todo) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TodoWhereInput) (*model.TodoConnection, error) {
	return u.todoRepository.List(ctx, after, first, before, last, where)
}

func (u *todo) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return u.todoRepository.Create(ctx, input)
}

func (u *todo) Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	return u.todoRepository.Update(ctx, input)
}


