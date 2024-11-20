package api

import (
	"io/ioutil"
	"log"
	"matchaVgo/internal/auth"
	"matchaVgo/internal/db"
	"matchaVgo/internal/schema"
	"matchaVgo/middleware"
	"net/http"
	"os"
    "path/filepath"
	"github.com/swaggo/http-swagger"
	_ "matchaVgo/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func parsedSchema() *graphql.Schema {

	schemaBytes, err := ioutil.ReadFile("internal/schema/schema.graphql");
	if err != nil {
		log.Fatalf("Failed to read schema file: %s", err);
	}
	schemaString := string(schemaBytes);
	return graphql.MustParseSchema(schemaString, &schema.Resolver{});

}

func graphqlHandler(schema *graphql.Schema) http.Handler {
	return &relay.Handler{Schema: schema}
}

// @title           matcher API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  matcherAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func Ga33ad_server() error {

	db := db.Connect();

	// Read and parse the schema
	parsedSchema := parsedSchema();

	// Set the database in the schema package
	schema.SetDB(db)
	router := mux.NewRouter();


	// Parent route
	apiRouter := router.PathPrefix("/api/v1").Subrouter();
	// Subrouter for swagger
	apiRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler);

	// Subrouter for auth
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.Handle("/login", graphqlHandler(parsedSchema)).Methods("POST");
	authRouter.Handle("/register", graphqlHandler(parsedSchema)).Methods("POST");
	authRouter.Handle("/send-verification-email", graphqlHandler(parsedSchema)).Methods("POST");
	authRouter.Handle("/reset-pass", graphqlHandler(parsedSchema)).Methods("POST");

	// Subrouter for users
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST");
	userRouter.Handle("/token", graphqlHandler(parsedSchema)).Methods("POST");
	userRouter.Handle("/proceed-registration", graphqlHandler(parsedSchema)).Methods("POST");
	userRouter.Handle("/update-info", graphqlHandler(parsedSchema)).Methods("POST");

	// Subrouter for posts
	postRouter := apiRouter.PathPrefix("/posts").Subrouter();
	postRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST");

	corsHandler := middleware.CORS(router);
	wrappedRouter := auth.WithResponseWriter(corsHandler);

	// Start the server
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	backendPort := os.Getenv("BACKEND_PORT");
	log.Println("✨ Running on port", backendPort, "..");
	return http.ListenAndServe(":"+backendPort, wrappedRouter);
}
