package repository

import (
	"context"
	"errors"
	"time"

	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/link"
	"go-basic-gqlgen/pkg/entity/model"
	usecaseRepository "go-basic-gqlgen/pkg/usecase/repository"
)

type linkRepository struct {
	client *ent.Client
}

func NewLinkRepository(client *ent.Client) usecaseRepository.Link {
	return &linkRepository{client: client}
}

func (r *linkRepository) Get(ctx context.Context, id *model.ID) (*model.Link, error) {
	u, err := r.client.Link.Query().Where(link.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, errors.New("link not found")
	}
	return u, nil
}

func (r *linkRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput) (*model.LinkConnection, error) {
	return r.client.Link.Query().Paginate(ctx, after, first, before, last, ent.WithLinkFilter(where.Filter))
}

func (r *linkRepository) Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	u, err := r.client.Link.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to create link")
	}
	return u, nil
}

func (r *linkRepository) Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	u, err := r.client.Link.UpdateOneID(input.ID).SetInput(input).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to update link")
	}
	return u, nil
}
