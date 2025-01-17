package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mkdirjava/graphql-todos/graph"
	"github.com/mkdirjava/graphql-todos/graph/model"

	// Replace username with your github username
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		ModelCache: model.NewModelCache(),
	}}))
	h.AddTransport(&transport.Websocket{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()

	r.Any("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
