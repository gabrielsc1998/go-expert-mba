package server

import (
	"fmt"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabrielsc1998/go-expert-mba/challenge-3/clean-architecture/internal/infra/graph"
)

type GraphQLServer struct {
	Port      string
	Resolvers interface{}
}

type GraphQLServerInput struct {
	Port      string
	Resolvers interface{}
}

func NewGraphQLServer(input GraphQLServerInput) *GraphQLServer {
	return &GraphQLServer{
		Port:      input.Port,
		Resolvers: input.Resolvers,
	}
}

func (s *GraphQLServer) Start() {
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: s.Resolvers.(graph.ResolverRoot)}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	fmt.Println("Starting GraphQL server on port", s.Port)
	http.ListenAndServe(":"+s.Port, nil)
}
