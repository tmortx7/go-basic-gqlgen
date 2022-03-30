package registry

import (
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/adapter/repository"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewEmployeeController() controller.Employee {
	repo := repository.NewEmployeeRepository(r.client)
	u := usecase.NewEmployeeUsecase(repo)

	return controller.NewEmployeeController(u)
}
