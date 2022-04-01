package controller

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

type group struct {
	groupUsecase usecase.Group
}

// Group of interface
type Group interface {
	Get(ctx context.Context, id *model.ID) (*model.Group, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.GroupWhereInput) (*model.GroupConnection, error)
	Create(ctx context.Context, input model.CreateGroupInput) (*model.Group, error)
	Update(ctx context.Context, input model.UpdateGroupInput) (*model.Group, error)
}

// NewGroupController returns group controller
func NewGroupController(uu usecase.Group) Group {
	return &group{groupUsecase: uu}
}

func (u *group) Get(ctx context.Context, id *model.ID) (*model.Group, error) {
	return u.groupUsecase.Get(ctx, id)
}

func (u *group) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.GroupWhereInput) (*model.GroupConnection, error) {
	return u.groupUsecase.List(ctx, after, first, before, last, where)
}

func (u *group) Create(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	return u.groupUsecase.Create(ctx, input)
}

func (u *group) Update(ctx context.Context, input model.UpdateGroupInput) (*model.Group, error) {
	return u.groupUsecase.Update(ctx, input)
}
