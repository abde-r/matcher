package api

import (
	"database/sql"
	"log"
	"matchaVgo/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr	string
	db		*sql.DB
}

type ApiError struct {
	Error string
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	
	// ** Registering services here **
	userStore := user.NewStore(s.db)
	userService := user.NewService(userStore)
	userService.RegisterRoutes(subrouter)


	// log.Fatal(http.ListenAndServe(s.addr, subrouter))
	log.Println("Listining on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
