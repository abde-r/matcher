package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type PostService struct {
	store Store
}

func NewPostService(s Store) *PostService {
	return &PostService{store: s}
}

func (s *PostService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/posts", s.handleGetPost).Method("GET")
	r.HandleFunc("/posts", s.handleCreatePost).Method("POST")
}

func (s *PostService) handleGetPost(w http.ResponseWriter, r *http.Request) {

}

func (s *PostService) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	
}