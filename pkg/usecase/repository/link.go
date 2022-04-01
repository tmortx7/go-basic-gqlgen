package repository

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
)

type Link interface {
	Get(ctx context.Context, id *model.ID) (*model.Link, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput) (*model.LinkConnection, error)
	Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error)
	Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error)
}
