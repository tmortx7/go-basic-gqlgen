package repository

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
)

type Todo interface {
	Get(ctx context.Context, id *model.ID) (*model.Todo, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TodoWhereInput) (*model.TodoConnection, error)
	Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error)
	Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error)
}
