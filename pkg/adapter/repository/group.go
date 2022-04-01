package repository

import (
	"context"
	"errors"

	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/group"
	"go-basic-gqlgen/pkg/entity/model"
	usecaseRepository "go-basic-gqlgen/pkg/usecase/repository"
)

type groupRepository struct {
	client *ent.Client
}

func NewGroupRepository(client *ent.Client) usecaseRepository.Group {
	return &groupRepository{client: client}
}

func (r *groupRepository) Get(ctx context.Context, id *model.ID) (*model.Group, error) {
	u, err := r.client.Group.Query().Where(group.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, errors.New("group not found")
	}
	return u, nil
}

func (r *groupRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.GroupWhereInput) (*model.GroupConnection, error) {
	return r.client.Group.Query().Paginate(ctx, after, first, before, last, ent.WithGroupFilter(where.Filter))
}

func (r *groupRepository) Create(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	u, err := r.client.Group.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to create group")
	}
	return u, nil
}

func (r *groupRepository) Update(ctx context.Context, input model.UpdateGroupInput) (*model.Group, error) {
	u, err := r.client.Group.UpdateOneID(input.ID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to update group")
	}

	return u, nil
}
