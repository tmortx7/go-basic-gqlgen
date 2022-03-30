package repository

import (
	"context"
	"errors"
	"time"

	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/ent/employee"
	"go-basic-gqlgen/pkg/entity/model"
	usecaseRepository "go-basic-gqlgen/pkg/usecase/repository"
)

type employeeRepository struct {
	client *ent.Client
}

func NewEmployeeRepository(client *ent.Client) usecaseRepository.Employee {
	return &employeeRepository{client: client}
}

func (r *employeeRepository) Get(ctx context.Context, id *model.ID) (*model.Employee, error) {
	u, err := r.client.Employee.Query().Where(employee.IDEQ(*id)).Only(ctx)
	if err != nil {
		return nil, errors.New("employee not found")
	}
	return u, nil
}

func (r *employeeRepository) List(ctx context.Context, after *model.Cursor, first *int, before *model.Cursor, last *int, where *model.EmployeeWhereInput) (*model.EmployeeConnection, error) {
	return r.client.Employee.Query().Paginate(ctx, after, first, before, last, ent.WithEmployeeFilter(where.Filter))
}

func (r *employeeRepository) Create(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	u, err := r.client.Employee.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to create employee")
	}
	return u, nil
}

func (r *employeeRepository) Update(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	u, err := r.client.Employee.UpdateOneID(input.ID).SetInput(input).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, errors.New("failed to update employee")
	}
	return u, nil
}
