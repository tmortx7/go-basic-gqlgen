package controller

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

type link struct {
	linkUsecase usecase.Link
}

// Link of interface
type Link interface {
	Get(ctx context.Context, id *model.ID) (*model.Link, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput) (*model.LinkConnection, error)
	Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error)
	Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error)
}

// NewLinkController returns link controller
func NewLinkController(uu usecase.Link) Link {
	return &link{linkUsecase: uu}
}

func (u *link) Get(ctx context.Context, id *model.ID) (*model.Link, error) {
	return u.linkUsecase.Get(ctx, id)
}

func (u *link) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput) (*model.LinkConnection, error) {
	return u.linkUsecase.List(ctx, after, first, before, last, where)
}

func (u *link) Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	return u.linkUsecase.Create(ctx, input)
}

func (u *link) Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	return u.linkUsecase.Update(ctx, input)
}
