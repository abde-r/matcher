package middleware

import (
	"context"
	"encoding/json"
	"fmt"
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



// func CustomHTTPMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		wrappedWriter := &responseWriterWrapper{ResponseWriter: w}
// 		next.ServeHTTP(wrappedWriter, r)

// 		// Adjust HTTP status code based on GraphQL errors
// 		if gqlErrs, ok := getErrorsFromContext(r.Context()); ok {
// 			for _, gqlErr := range gqlErrs {
// 				if gqlErr.Extensions != nil {
// 					if statusCode, exists := gqlErr.Extensions["statusCode"].(int); exists {
// 						wrappedWriter.WriteHeader(statusCode)
// 						return
// 					}
// 				}
// 			}
// 		}
// 	})
// }

// type contextKey string

// const errorContextKey contextKey = "errors"

// func getErrorsFromContext(ctx context.Context) ([]string, bool) {
// 	// Try to get the errors from context
// 	if err, ok := ctx.Value(errorContextKey).([]string); ok {
// 		return err, true
// 	}
// 	return nil, false
// }

func CustomHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Wrap the ResponseWriter to intercept the status code
		wrappedWriter := &responseWriterWrapper{ResponseWriter: w}

		// Call the next handler in the chain
		next.ServeHTTP(wrappedWriter, r)

		// Check for errors in the response context
		if wrappedWriter.statusCode == 0 {
			// Default status code in case no status is set
			wrappedWriter.statusCode = http.StatusOK
		}

		// Check if the response contains GraphQL errors (you may need to inspect the response body)
		if wrappedWriter.statusCode == http.StatusUnauthorized {
			// In this case, set the status code for the error explicitly
			// If the response contains errors, you can enrich it with your custom logic
			if err := wrappedWriter.setErrorCode(r.Context(), http.StatusUnauthorized); err != nil {
				fmt.Printf("Error while setting error code: %v", err)
			}
		}

		// Log the final status code for debugging
		fmt.Printf("Final Response Status Code: %d\n", wrappedWriter.statusCode)

		// Write the final status code to the ResponseWriter
		w.WriteHeader(wrappedWriter.statusCode)
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
	body       []byte // Capture the response body
}

func (rw *responseWriterWrapper) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriterWrapper) Write(p []byte) (n int, err error) {
	// Capture the body of the response
	rw.body = append(rw.body, p...)
	return rw.ResponseWriter.Write(p)
}

// Method to enrich the error response
func (rw *responseWriterWrapper) setErrorCode(_ context.Context, code int) error {
	// Example: Set the status code on errors
	if rw.body != nil && len(rw.body) > 0 {
		// You can parse the response body and add the `status_code` to errors here
		// For simplicity, let's assume we're modifying the GraphQL error response

		var graphQLResponse map[string]interface{}
		if err := json.Unmarshal(rw.body, &graphQLResponse); err != nil {
			return fmt.Errorf("failed to unmarshal response body: %v", err)
		}

		if errors, exists := graphQLResponse["errors"]; exists {
			// Assuming errors are of type []interface{}
			for _, errItem := range errors.([]interface{}) {
				if errMap, ok := errItem.(map[string]interface{}); ok {
					errMap["status_code"] = code // Add your custom status code here
				}
			}
		}

		// Re-encode the modified response back into the body
		updatedBody, err := json.Marshal(graphQLResponse)
		if err != nil {
			return fmt.Errorf("failed to marshal updated response: %v", err)
		}

		// Write the modified body back to the ResponseWriter
		rw.body = updatedBody
	}

	return nil
}

