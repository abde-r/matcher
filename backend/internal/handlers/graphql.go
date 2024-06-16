package handlers

import (
	"io/ioutil"
	"log"
	"matchaVgo/internal/schema"
	// "net/http"

	"github.com/graph-gophers/graphql-go"
	// "github.com/graph-gophers/graphql-go/relay"
	"github.com/jmoiron/sqlx"
)

func GraphQLHandler(db *sqlx.DB) *graphql.Schema {

	schema.SetDB(db)

	// reading and parsing schema
	schemaBytes, err := ioutil.ReadFile("internal/schema/schema.graphql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %s", err)
	}
	schemaString := string(schemaBytes)
	// parsedSchema := graphql.MustParseSchema(schemaString, &schema.Resolver{})

	return graphql.MustParseSchema(schemaString, &schema.Resolver{});
}