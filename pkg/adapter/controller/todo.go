package controller

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

type todo struct {
	todoUsecase usecase.Todo
}

// Todo of interface
type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TodoWhereInput) (*model.TodoConnection, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
	Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error)
}

// NewTodoController returns todo controller
func NewTodoController(uu usecase.Todo) Todo {
	return &todo{todoUsecase: uu}
}

func (u *todo) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	return u.todoUsecase.Get(ctx, id)
}

func (u *todo) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TodoWhereInput) (*model.TodoConnection, error) {
	return u.todoUsecase.List(ctx, after, first, before, last, where)
}

func (u *todo) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	return u.todoUsecase.Create(ctx, input)
}

func (u *todo) Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	return u.todoUsecase.Update(ctx, input)
}
