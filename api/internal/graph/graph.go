package graph

import (
	"net/http"

	"api/internal/graph/dataloader"
	"api/internal/graph/generated"
	"api/internal/graph/resolver"
	"api/internal/service"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

//go:generate go run github.com/99designs/gqlgen generate ./gqlgen.yml
//go:generate gofumpt -w ./resolver

func New(service service.Service, dataloader *dataloader.Dataloader) http.Handler {
	c := generated.Config{
		Resolvers: resolver.New(service, dataloader),
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return srv
}

func Playground(queryRoute string) http.Handler {
	return playground.Handler("GraphQL playground", queryRoute)
}
