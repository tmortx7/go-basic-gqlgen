package repository

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
)

type Employee interface {
	Get(ctx context.Context, id *model.ID) (*model.Employee, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.EmployeeWhereInput) (*model.EmployeeConnection, error)
	Create(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error)
	Update(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error)
}
