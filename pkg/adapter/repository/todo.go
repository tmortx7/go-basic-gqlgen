package repository

import (
	"context"
	"errors"

	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/todo"
	"go-basic-gqlgen/pkg/entity/model"
	usecaseRepository "go-basic-gqlgen/pkg/usecase/repository"
)

type todoRepository struct {
	client *ent.Client
}

func NewTodoRepository(client *ent.Client) usecaseRepository.Todo {
	return &todoRepository{client: client}
}

func (r *todoRepository) Get(ctx context.Context, id *model.ID) (*model.Todo, error) {
	u, err := r.client.Todo.Query().Where(todo.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	return u, nil
}

func (r *todoRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.TodoWhereInput) (*model.TodoConnection, error) {
	return r.client.Todo.Query().Paginate(ctx, after, first, before, last, ent.WithTodoFilter(where.Filter))
}

func (r *todoRepository) Create(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	u, err := r.client.Todo.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to create todo")
	}
	return u, nil
}

func (r *todoRepository) Update(ctx context.Context, input model.UpdateTodoInput) (*model.Todo, error) {
	u, err := r.client.Todo.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to update todo")
	}
	return u, nil
}
