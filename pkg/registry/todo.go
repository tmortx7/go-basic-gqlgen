package registry

import (
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/adapter/repository"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewTodoController() controller.Todo {
	repo := repository.NewTodoRepository(r.client)
	u := usecase.NewTodoUsecase(repo)

	return controller.NewTodoController(u)
}
