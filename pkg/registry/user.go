package registry

import (
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/adapter/repository"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewUserController() controller.User {
	repo := repository.NewUserRepository(r.client)
	u := usecase.NewUserUsecase(repo)

	return controller.NewUserController(u)
}
