package schema

import (
	// "time"

	// "github.com/graph-gophers/graphql-go"
)

// Resolver struct
type Resolver struct{}

type ProceedRegisterationUserInput struct {
	ID        	int32
	First_name	string
	Last_name	string
	Birthday	string
	Gender		bool
	Preferences	string
	Pics		string
	Location 	string
}