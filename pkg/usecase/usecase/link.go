package usecase

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/repository"
)

type link struct {
	linkRepository repository.Link
}

// Link of linkcase
type Link interface {
	Get(ctx context.Context, id *model.ID) (*model.Link, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput) (*model.LinkConnection, error)
	Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error)
	Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error)
}

// NewLinkUsecase returns link usecse
func NewLinkUsecase(r repository.Link) Link {
	return &link{linkRepository: r}
}

func (u *link) Get(ctx context.Context, id *model.ID) (*model.Link, error) {
	return u.linkRepository.Get(ctx, id)
}

func (u *link) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.LinkWhereInput) (*model.LinkConnection, error) {
	return u.linkRepository.List(ctx, after, first, before, last, where)
}

func (u *link) Create(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	return u.linkRepository.Create(ctx, input)
}

func (u *link) Update(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	return u.linkRepository.Update(ctx, input)
}


