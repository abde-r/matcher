package main

import (
	"encoding/json"
	"io"
	"log"
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
	r.HandleFunc("/posts", s.handleGetAllPost).Methods("GET")
	r.HandleFunc("/posts/{id}", s.handleGetPost).Methods("GET")
	r.HandleFunc("/posts", s.handleCreatePost).Methods("POST")
}

func (s *PostService) handleGetAllPost(w http.ResponseWriter, r *http.Request) {

	posts, err :=  s.store.GetAllPosts()//[...]string{"post1", "post2", "post3"}
	log.Println("post", posts)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error getting Project"})
		return
	}
	WriteJSON(w, http.StatusOK, posts)
}

func (s *PostService) handleGetPost(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	post, err :=  s.store.GetPostById(id)//[...]string{"post1", "post2", "post3"}
	log.Println("post", post)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error getting Project"})
		return
	}
	WriteJSON(w, http.StatusOK, post)
}

func (s *PostService) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var post *Post
	err = json.Unmarshal(body, &post)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if post.Title == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Title is required"})
		return
	}

	err = s.store.CreatePost(post)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating project"})
		return
	}

	WriteJSON(w, http.StatusCreated, post)
}