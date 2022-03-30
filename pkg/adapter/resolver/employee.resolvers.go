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

func (r *employeeResolver) CreatedAt(ctx context.Context, obj *ent.Employee) (string, error) {
	return datetime.FormatDate(obj.CreatedAt), nil
}

func (r *employeeResolver) UpdatedAt(ctx context.Context, obj *ent.Employee) (string, error) {
	return datetime.FormatDate(obj.UpdatedAt), nil
}

func (r *mutationResolver) CreateEmployee(ctx context.Context, input ent.CreateEmployeeInput) (*ent.Employee, error) {
	return r.controller.Employee.Create(ctx, input)
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, input ent.UpdateEmployeeInput) (*ent.Employee, error) {
	return r.controller.Employee.Update(ctx, input)
}

func (r *queryResolver) SingleEmployee(ctx context.Context, id *ulid.ID) (*ent.Employee, error) {
	return r.controller.Employee.Get(ctx, id)
}

func (r *queryResolver) AllEmployees(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.EmployeeWhereInput) (*ent.EmployeeConnection, error) {
	return r.controller.Employee.List(ctx, after, first, before, last, where)
}

// Employee returns generated.EmployeeResolver implementation.
func (r *Resolver) Employee() generated.EmployeeResolver { return &employeeResolver{r} }

type employeeResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Employee(ctx context.Context, id *ulid.ID) (*ent.Employee, error) {
	return r.controller.Employee.Get(ctx, id)
}
func (r *queryResolver) Employees(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.EmployeeWhereInput) (*ent.EmployeeConnection, error) {
	return r.controller.Employee.List(ctx, after, first, before, last, where)
}
