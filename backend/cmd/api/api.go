package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"matchaVgo/internal/schema"
	"strings"
	// "time"

	// "matchaVgo/middleware"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jmoiron/sqlx"
	// "github.com/lib/pq"
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
	// router.Use(middleware.CORS);

	// Parent route
	apiRouter := router.PathPrefix("/api/v1").Subrouter();

	// Subrouter for auth
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	authRouter.Handle("/login", graphqlHandler(parsedSchema)).Methods("POST");
	authRouter.Handle("/register", graphqlHandler(parsedSchema)).Methods("POST");

	// Subrouter for users
	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST");
	// userRouter.Handle("/g", graphqlHandler(parsedSchema)).Methods("GET");
	userRouter.Handle("/proceed-registration", graphqlHandler(parsedSchema)).Methods("POST");

	// Subrouter for posts
	postRouter := apiRouter.PathPrefix("/posts").Subrouter();
	postRouter.Handle("/", graphqlHandler(parsedSchema)).Methods("POST");

	corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow only your frontend origin for security
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
        handlers.AllowCredentials(),
    )(router)

	// Start the server
	log.Println("✨ Running on port", backendPort, "..");
	return http.ListenAndServe(":"+backendPort, corsHandler);
}

func graphqlHandler(schema *graphql.Schema) http.Handler {
	return &relay.Handler{Schema: schema}
}

func proceedRegistrationHandler(schema *graphql.Schema) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestBody struct {
			Input struct {
				ID			int32		`json:"id"`
				FirstName   string    	`json:"first_name"`
				LastName    string    	`json:"last_name"`
				Birthday 	string    	`json:"birthday"` // Use string to parse from JSON
				Gender      bool      	`json:"gender"`
				Preferences []string  	`json:"preferences"`
				Pics 		[]string  	`json:"pics"`
				Location	string  	`json:"location"`
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

		// Parse the birthday string to time.Time
		// birthday, err := time.Parse(time.RFC3339, requestBody.Input.BirthdayStr)
		// if err != nil {
		// 	http.Error(w, "Invalid birthday format", http.StatusBadRequest)
		// 	return
		// }

		// Insert user into the database
		// var userID int
		// query := `INSERT INTO users (first_name, last_name, birthday, gender, preferences) VALUES ($1, $2, $3, $4, $5) RETURNING id`
		// err = db.QueryRow(query, requestBody.Input.FirstName, requestBody.Input.LastName, birthday, requestBody.Input.Gender, pq.Array(requestBody.Input.Preferences)).Scan(&requestBody.Input.ID)
		// if err != nil {
		// 	http.Error(w, fmt.Sprintf("Failed to insert user: %v", err), http.StatusInternalServerError)
		// 	return
		// }
		prefs := "\""+strings.Join(requestBody.Input.Preferences, ";")+"\"";
		pics := "\""+strings.Join(requestBody.Input.Pics, ";;;")+"\"";
		
		query := fmt.Sprintf(`mutation { proceedRegistrationUser(input: {id: %d, first_name: "%s", last_name: "%s", birthday: "%s", gender: %t, preferences: %s , pics: %s, location: "%s"}) { id, first_name, last_name, birthday, gender, preferences, pics, location } }`,
			requestBody.Input.ID, requestBody.Input.FirstName, requestBody.Input.LastName, requestBody.Input.Birthday, requestBody.Input.Gender, prefs, pics, requestBody.Input.Location)
		reqBody := fmt.Sprintf(`{"query": %q}`, query)
		
		newReq, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBufferString(reqBody))
		if err != nil {
			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
			return
		}
		newReq.Header.Set("Content-Type", "application/json")
		
		// fmt.Println("hola", query)
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
