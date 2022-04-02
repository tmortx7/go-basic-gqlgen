package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/schema/ulid"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	return r.controller.Todo.Create(ctx, input)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input ent.UpdateTodoInput) (*ent.Todo, error) {
	return r.controller.Todo.Update(ctx, input)
}

func (r *queryResolver) SingleTodo(ctx context.Context, id *ulid.ID) (*ent.Todo, error) {
	return r.controller.Todo.Get(ctx, id)
}

func (r *queryResolver) AllTodos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.controller.Todo.List(ctx, after, first, before, last, where)
}
