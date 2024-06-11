package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errFirstNameRequired = errors.New("firstname is required")
var errLastNameRequired = errors.New("lastname is required")
var errUsernameRequired = errors.New("username is required")
var errEmailRequired = errors.New("email is required")
var errPasswordRequired = errors.New("password is required")
var errGenderRequired = errors.New("gender is required")

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/user/register", s.handleUserRegister).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var payload *User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{ Error: "Invalid request payload" })
		return
	}

	if err := validateUserPayload(payload); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{ Error: err.Error() })
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}
	payload.Password = hashedPassword

	u, err := s.store.CreateUser(payload)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{ Error: "Error creating user" })
		return
	}

	token, err := createAndSetAuthCookie(u.Id, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{ Error: "Error creating session" })
		return
	}

	WriteJSON(w, http.StatusCreated, token)

}

func validateUserPayload(user *User) error {
	if user.FirstName == "" {
		return errFirstNameRequired
	}
	if user.LastName == "" {
		return errLastNameRequired
	}
	if user.Username == "" {
		return errUsernameRequired
	}
	if user.Email == "" {
		return errEmailRequired
	}
	if user.Password == "" {
		return errPasswordRequired
	}
	if user.Gender == "" {
		return errGenderRequired
	}
	return nil
}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		return "", nil
	}

	http.SetCookie(w, &http.Cookie{
		Name: "Authorization",
		Value: token,
	})

	return token, nil
}