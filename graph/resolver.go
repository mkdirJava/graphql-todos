package graph

import "github.com/mkdirjava/graphql-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ModelCache model.ModelCacher
}
