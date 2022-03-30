package resolver

import (
	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/graph/generated"
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/adapter/directives"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	}
	c.Directives.Binding = directives.Binding

	return generated.NewExecutableSchema(c)
}
