package registry

import (
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/adapter/repository"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewGroupController() controller.Group {
	repo := repository.NewGroupRepository(r.client)
	u := usecase.NewGroupUsecase(repo)

	return controller.NewGroupController(u)
}
