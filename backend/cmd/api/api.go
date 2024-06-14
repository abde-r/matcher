package api

import (
	// "database/sql"
	"log"
	"matchaVgo/services/user"
	"net/http"
	// "github.com/graphql-go/handler"
	// "matchaVgo/internal/schema"



	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type APIServer struct {
	addr	string
	db		*sqlx.DB
}

type ApiError struct {
	Error string
}

func NewAPIServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	// /// Graphql handler
	// _handler := handler.New(&handler.Config{
	// 	Schema: &schema.Schema,
	// 	Pretty: true,
	// 	GraphiQL: true,
	// });
	// http.Handle("/graphql", _handler);
	// // http.ListenAndServe("8000", nil);
	// log.Print("Running on", "8000");

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
