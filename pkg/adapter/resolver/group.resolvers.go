package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	user "go-basic-gqlgen"
	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/schema/ulid"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*ent.Group, error) {
	return r.controller.Group.Create(ctx, input)
}

func (r *mutationResolver) UpdateGroup(ctx context.Context, input ent.UpdateGroupInput) (*ent.Group, error) {
	return r.controller.Group.Update(ctx, input)
}

func (r *mutationResolver) CreateGroupUser(ctx context.Context, input user.GroupUserInput) (*ent.Group, error) {
	return ent.FromContext(ctx).Group.UpdateOneID(input.GroupID).
		AddUserIDs(input.UserID).
		Save(ctx)
}

func (r *queryResolver) SingleGroup(ctx context.Context, id *ulid.ID) (*ent.Group, error) {
	return r.controller.Group.Get(ctx, id)
}

func (r *queryResolver) AllGroups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.controller.Group.List(ctx, after, first, before, last, where)
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Update(ctx context.Context, input ent.UpdateGroupInput) (*ent.Group, error) {
	return r.controller.Group.Update(ctx, input)
}
