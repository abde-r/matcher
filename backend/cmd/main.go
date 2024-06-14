package main

import (
	"log"
	"matchaVgo/internal/db"
	"matchaVgo/internal/handlers"
	"net/http"
	"os"
)

// func dbHandler() *sqlx.DB {

// 	err := godotenv.Load();
// 	if err != nil {
// 		log.Fatal("Error loading .env file");
// 	}

// 	db_name := os.Getenv("DB_NAME");
// 	db_user := os.Getenv("DB_USER");
// 	db_password := os.Getenv("DB_PASSWORD");

// 	s := "user=" + db_user + " password=" + db_password + " dbname=" + db_name + " sslmode=disable";
// 	db, err := sqlx.Connect("postgres", s);
// 	if err != nil {
// 		log.Fatalln(err);
// 	}

// 	return db;
// }

// func graphqlHandler(db *sqlx.DB) http.Handler {

// 	schema.SetDB(db)

// 	// reading and parsing schema
// 	schemaBytes, err := ioutil.ReadFile("internal/schema/schema.graphql")
// 	if err != nil {
// 		log.Fatalf("Failed to read schema file: %s", err)
// 	}
// 	schemaString := string(schemaBytes)
// 	parsedSchema := graphql.MustParseSchema(schemaString, &schema.Resolver{})

// 	// Graphql handler
// 	return &relay.Handler{Schema: parsedSchema}

// }

func main() {
	
	db := db.Connect();
	_graphqlHandler := handlers.GraphQLHandler(db);

	http.Handle("/api/v1", _graphqlHandler);
	backend_port := os.Getenv("BACKEND_PORT");

	log.Print("Running on port ", backend_port);
	if err := http.ListenAndServe(":"+backend_port, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err);
	}
}
