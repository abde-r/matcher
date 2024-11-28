package api

import (
	"log"
	_ "matcher/docs"
	"matcher/internal/auth"
	"matcher/internal/db"
	"matcher/internal/schema"
	"matcher/internal/store"
	"matcher/middleware"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func parsedSchema() *graphql.Schema {
	// Read the schema file
	schemaBytes, err := os.ReadFile("internal/schema/schema.graphql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %s", err)
	}
	schemaString := string(schemaBytes)

	// Parse schema and return schema object
	return graphql.MustParseSchema(schemaString, &schema.Resolver{})
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

	db := db.Connect()

	// Read and parse the schema
	parsedSchema := parsedSchema()

	// Set the database in the schema package
	schema.SetDB(db)
	router := mux.NewRouter()

	// Run hub for sockets
	go store.HubRunner()

	// Parent route
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	// Subrouter for swagger
	apiRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Subrouter for auth
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	// authRouter.Handle("/login", graphqlHandler(parsedSchema)).Methods("POST")
	authRouter.Handle("/login", middleware.CustomHTTPMiddleware(graphqlHandler(parsedSchema))).Methods("POST")
	authRouter.Handle("/register", graphqlHandler(parsedSchema)).Methods("POST")
	authRouter.Handle("/send-verification-email", graphqlHandler(parsedSchema)).Methods("POST")
	authRouter.Handle("/account-verification", graphqlHandler(parsedSchema)).Methods("POST")
	authRouter.Handle("/reset-pass", graphqlHandler(parsedSchema)).Methods("POST")

	// Subrouter for users
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST")
	userRouter.Handle("/token", graphqlHandler(parsedSchema)).Methods("POST")
	userRouter.Handle("/proceed-registration", graphqlHandler(parsedSchema)).Methods("POST")
	userRouter.Handle("/update-info", graphqlHandler(parsedSchema)).Methods("POST")

	// Subrouter for posts
	postRouter := apiRouter.PathPrefix("/posts").Subrouter()
	postRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST")

	// Subrouter for chat
	apiRouter.HandleFunc("/chat", store.WebSocketHandler)
	// chatRouter := apiRouter.PathPrefix("/chat").Subrouter();
	// chatRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST");

	corsHandler := middleware.CORS(router)
	wrappedRouter := auth.WithResponseWriter(corsHandler)

	// Start the server
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	backendPort := os.Getenv("BACKEND_PORT")

	log.Println("✨ Running on port", backendPort, "..")
	return http.ListenAndServe(":"+backendPort, wrappedRouter)
}
