package middleware

import (
	"encoding/json"
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

type customResponseWriter struct {
    http.ResponseWriter
    statusCode int
    body       []byte
}

func (c *customResponseWriter) WriteHeader(statusCode int) {
    c.statusCode = statusCode
    c.ResponseWriter.WriteHeader(statusCode)
}

func (c *customResponseWriter) Write(data []byte) (int, error) {
    c.body = append(c.body, data...)
    return c.ResponseWriter.Write(data)
}

func CustomHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customWriter := &customResponseWriter{ResponseWriter: w}
		next.ServeHTTP(customWriter, r)

		var response struct {
			Data   interface{}              `json:"data"`
			Errors []map[string]interface{} `json:"errors"`
		}

		// Parse the response body
		if err := json.Unmarshal(customWriter.body, &response); err == nil && len(response.Errors) > 0 {
			for _, err := range response.Errors {
				// Add statusCode only if it's not already present
				if extensions, ok := err["extensions"].(map[string]interface{}); ok {
					if _, hasStatusCode := extensions["statusCode"]; !hasStatusCode {
						if code, ok := extensions["code"].(float64); ok {
							extensions["statusCode"] = int(code)
						} else {
							extensions["statusCode"] = http.StatusInternalServerError
						}
					}
				} else {
					err["extensions"] = map[string]interface{}{
						"statusCode": http.StatusInternalServerError,
					}
				}
			}

			// Re-marshal and overwrite the response
			customWriter.body, _ = json.Marshal(response)
			customWriter.ResponseWriter.Header().Set("Content-Type", "application/json")
			customWriter.ResponseWriter.WriteHeader(customWriter.statusCode)
			customWriter.ResponseWriter.Write(customWriter.body)
		}
	})
}




// type responseWriterWrapper struct {
// 	http.ResponseWriter
// 	statusCode int
// 	body       []byte // Capture the response body
// }

// func (rw *responseWriterWrapper) WriteHeader(code int) {
// 	rw.statusCode = code
// 	rw.ResponseWriter.WriteHeader(code)
// }

// func (rw *responseWriterWrapper) Write(p []byte) (n int, err error) {
// 	// Capture the body of the response
// 	rw.body = append(rw.body, p...)
// 	return rw.ResponseWriter.Write(p)
// }

// // Method to enrich the error response
// func (rw *responseWriterWrapper) setErrorCode(_ context.Context, code int) error {
// 	// Example: Set the status code on errors
// 	if rw.body != nil && len(rw.body) > 0 {
// 		// You can parse the response body and add the `status_code` to errors here
// 		// For simplicity, let's assume we're modifying the GraphQL error response

// 		var graphQLResponse map[string]interface{}
// 		if err := json.Unmarshal(rw.body, &graphQLResponse); err != nil {
// 			return fmt.Errorf("failed to unmarshal response body: %v", err)
// 		}

// 		if errors, exists := graphQLResponse["errors"]; exists {
// 			// Assuming errors are of type []interface{}
// 			for _, errItem := range errors.([]interface{}) {
// 				if errMap, ok := errItem.(map[string]interface{}); ok {
// 					errMap["status_code"] = code // Add your custom status code here
// 				}
// 			}
// 		}

// 		// Re-encode the modified response back into the body
// 		updatedBody, err := json.Marshal(graphQLResponse)
// 		if err != nil {
// 			return fmt.Errorf("failed to marshal updated response: %v", err)
// 		}

// 		// Write the modified body back to the ResponseWriter
// 		rw.body = updatedBody
// 	}

// 	return nil
// }

