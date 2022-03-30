package controller

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

type employee struct {
	employeeUsecase usecase.Employee
}

// Employee of interface
type Employee interface {
	Get(ctx context.Context, id *model.ID) (*model.Employee, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.EmployeeWhereInput) (*model.EmployeeConnection, error)
	Create(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error)
	Update(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error)
}

// NewEmployeeController returns employee controller
func NewEmployeeController(uu usecase.Employee) Employee {
	return &employee{employeeUsecase: uu}
}

func (u *employee) Get(ctx context.Context, id *model.ID) (*model.Employee, error) {
	return u.employeeUsecase.Get(ctx, id)
}

func (u *employee) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.EmployeeWhereInput) (*model.EmployeeConnection, error) {
	return u.employeeUsecase.List(ctx, after, first, before, last, where)
}

func (u *employee) Create(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	return u.employeeUsecase.Create(ctx, input)
}

func (u *employee) Update(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	return u.employeeUsecase.Update(ctx, input)
}
