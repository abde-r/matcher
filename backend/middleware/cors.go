package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CORS middleware function
func CORS(router *mux.Router) http.Handler {

	back_url := os.Getenv("FRONT_URL");
	return handlers.CORS(
		    handlers.AllowedOrigins([]string{back_url}),
		    handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		    handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		    handlers.AllowCredentials(),
		)(router)
}
