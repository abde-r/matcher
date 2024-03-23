package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello 'go' matcha")

	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello matchaV-go")
	})
}
