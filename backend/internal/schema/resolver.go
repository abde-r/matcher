package schema

import "github.com/graph-gophers/graphql-go"

// Resolver struct
type Resolver struct{}

type CompleteRegisterationUserInput struct {
	ID        graphql.ID
	First_name string
	Last_name  string
	Gender    bool
}