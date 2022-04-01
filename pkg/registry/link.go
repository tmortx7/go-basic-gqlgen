package registry

import (
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/adapter/repository"
	"go-basic-gqlgen/pkg/usecase/usecase"
)

func (r *registry) NewLinkController() controller.Link {
	repo := repository.NewLinkRepository(r.client)
	u := usecase.NewLinkUsecase(repo)

	return controller.NewLinkController(u)
}
