package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "matchaVgo/internal/handlers"
	"github.com/gorilla/handlers"
	"matchaVgo/internal/schema"
	"matchaVgo/middleware"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	// "github.com/graphql-go/handler"
	"github.com/jmoiron/sqlx"
)

func Ga33ad_server(db *sqlx.DB) error {
	// Read and parse the schema
	schemaBytes, err := ioutil.ReadFile("internal/schema/schema.graphql")
	if err != nil {
		log.Fatalf("Failed to read schema file: %s", err)
	}
	schemaString := string(schemaBytes)
	parsedSchema := graphql.MustParseSchema(schemaString, &schema.Resolver{})

	// Set the database in the schema package
	schema.SetDB(db)

	// Get backend port from environment variable
	backendPort := os.Getenv("BACKEND_PORT");
	if backendPort == "" {
		backendPort = "8000";
	}

	// Create a new router
	router := mux.NewRouter();
	router.Use(middleware.CORS);

	// Parent route
	apiRouter := router.PathPrefix("/api/v1").Subrouter();

	// Subrouter for auth
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.Handle("/login", graphqlHandler(parsedSchema)).Methods("POST");
	authRouter.Handle("/register", graphqlHandler(parsedSchema)).Methods("POST");

	// Subrouter for users
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.Handle("/g", graphqlHandler(parsedSchema)).Methods("GET");
	// userRouter.Handle("/g", graphqlHandler(parsedSchema)).Methods("GET");
	userRouter.Handle("/complete-registration", completeRegistrationHandler(parsedSchema)).Methods("POST");

	// Subrouter for posts
	postRouter := apiRouter.PathPrefix("/posts").Subrouter();
	postRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST");

	// Start the server
	log.Println("✨ Running on port", backendPort, "..");
	// return http.ListenAndServe(":"+backendPort,handlers.CORS(
	// 		handlers.AllowedOrigins([]string{"*"}),
	// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	)(router));
	
	return http.ListenAndServe(":"+backendPort, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(router))
}

func graphqlHandler(schema *graphql.Schema) http.Handler {
	return &relay.Handler{Schema: schema}
}

func completeRegistrationHandler(schema *graphql.Schema) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Input struct {
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Gender    bool   `json:"gender"`
			} `json:"input"`
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &requestBody)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		query := fmt.Sprintf(`mutation { completeRegistration(input: {first_name: "%s", last_name: "%s", gender: %t}) { id, first_name, last_name, gender } }`,
			requestBody.Input.FirstName, requestBody.Input.LastName, requestBody.Input.Gender)
		reqBody := fmt.Sprintf(`{"query": %q}`, query) // Properly escape the query

		newReq, err := http.NewRequest("POST", "/graphql", bytes.NewBufferString(reqBody))
		if err != nil {
			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
			return
		}
		newReq.Header.Set("Content-Type", "application/json")

		// Create a new response recorder to capture the response
		recorder := httptest.NewRecorder()
		h := &relay.Handler{Schema: schema}
		h.ServeHTTP(recorder, newReq)

		// Copy the recorded response to the original response writer
		for k, v := range recorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(recorder.Code)
		w.Write(recorder.Body.Bytes())
	})
}