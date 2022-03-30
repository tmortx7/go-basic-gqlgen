package main

import (
	"go-basic-gqlgen/config"
	"go-basic-gqlgen/ent"
	"go-basic-gqlgen/pkg/adapter/controller"
	"go-basic-gqlgen/pkg/infrastructure/datastore"
	"go-basic-gqlgen/pkg/infrastructure/graphql"
	"go-basic-gqlgen/pkg/infrastructure/router"
	"go-basic-gqlgen/pkg/registry"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	e := router.New(srv)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening postgres client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}
