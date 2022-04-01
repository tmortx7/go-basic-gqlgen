package repository

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
)

type Group interface {
	Get(ctx context.Context, id *model.ID) (*model.Group, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.GroupWhereInput) (*model.GroupConnection, error)
	Create(ctx context.Context, input model.CreateGroupInput) (*model.Group, error)
	Update(ctx context.Context, input model.UpdateGroupInput) (*model.Group, error)
}
