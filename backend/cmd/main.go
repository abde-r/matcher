package main

import (
	"log"
	"matchaVgo/cmd/api"
	"matchaVgo/internal/db"
	// "matchaVgo/internal/handlers"
)

func main() {
	
	db := db.Connect();
	// schema := handlers.GraphQLHandler(db);

	if err := api.Ga33ad_server(db); err != nil {
		log.Fatal(err);
	}
}
