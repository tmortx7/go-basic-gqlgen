package usecase

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/repository"
)

type group struct {
	groupRepository repository.Group
}

// Group of groupcase
type Group interface {
	Get(ctx context.Context, id *model.ID) (*model.Group, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.GroupWhereInput) (*model.GroupConnection, error)
	Create(ctx context.Context, input model.CreateGroupInput) (*model.Group, error)
	Update(ctx context.Context, input model.UpdateGroupInput) (*model.Group, error)
}

// NewGroupUsecase returns group usecse
func NewGroupUsecase(r repository.Group) Group {
	return &group{groupRepository: r}
}

func (u *group) Get(ctx context.Context, id *model.ID) (*model.Group, error) {
	return u.groupRepository.Get(ctx, id)
}

func (u *group) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.GroupWhereInput) (*model.GroupConnection, error) {
	return u.groupRepository.List(ctx, after, first, before, last, where)
}

func (u *group) Create(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	return u.groupRepository.Create(ctx, input)
}

func (u *group) Update(ctx context.Context, input model.UpdateGroupInput) (*model.Group, error) {
	return u.groupRepository.Update(ctx, input)
}


