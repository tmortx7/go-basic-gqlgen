package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/schema/ulid"
	"go-basic-gqlgen/graph/generated"
	"go-basic-gqlgen/pkg/util/datetime"
)

func (r *linkResolver) CreatedAt(ctx context.Context, obj *ent.Link) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *linkResolver) UpdatedAt(ctx context.Context, obj *ent.Link) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *mutationResolver) CreateLink(ctx context.Context, input ent.CreateLinkInput) (*ent.Link, error) {
	//return r.controller.Link.Create(ctx, input)
	client := ent.FromContext(ctx)
	return client.Link.
		Create().
		SetTitle(input.Title).
		SetAddress(input.Address).
		SetUserID("usr_01fzhpnnps7q5kvkt69afvw807").
		Save(ctx)
}

func (r *mutationResolver) UpdateLink(ctx context.Context, input ent.UpdateLinkInput) (*ent.Link, error) {
	return r.controller.Link.Update(ctx, input)
}

func (r *queryResolver) SingleLink(ctx context.Context, id *ulid.ID) (*ent.Link, error) {
	return r.controller.Link.Get(ctx, id)
}

func (r *queryResolver) AllLinks(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.LinkWhereInput) (*ent.LinkConnection, error) {
	return r.controller.Link.List(ctx, after, first, before, last, where)
}

// Link returns generated.LinkResolver implementation.
func (r *Resolver) Link() generated.LinkResolver { return &linkResolver{r} }

type linkResolver struct{ *Resolver }
