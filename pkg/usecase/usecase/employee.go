package usecase

import (
	"context"
	"go-basic-gqlgen/pkg/entity/model"
	"go-basic-gqlgen/pkg/usecase/repository"
)

type employee struct {
	employeeRepository repository.Employee
}

// Employee of employeecase
type Employee interface {
	Get(ctx context.Context, id *model.ID) (*model.Employee, error)
	List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.EmployeeWhereInput) (*model.EmployeeConnection, error)
	Create(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error)
	Update(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error)
}

// NewEmployeeUsecase returns employee usecse
func NewEmployeeUsecase(r repository.Employee) Employee {
	return &employee{employeeRepository: r}
}

func (u *employee) Get(ctx context.Context, id *model.ID) (*model.Employee, error) {
	return u.employeeRepository.Get(ctx, id)
}

func (u *employee) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.EmployeeWhereInput) (*model.EmployeeConnection, error) {
	return u.employeeRepository.List(ctx, after, first, before, last, where)
}

func (u *employee) Create(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	return u.employeeRepository.Create(ctx, input)
}

func (u *employee) Update(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	return u.employeeRepository.Update(ctx, input)
}


