package user

import (
	"fmt"
	"matchaVgo/configs"
	"matchaVgo/services/auth"
	"matchaVgo/types"
	"matchaVgo/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type UserService struct {
	store types.UserStore
}

func NewService(store types.UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", s.UserLogin).Methods(http.MethodPost)
	r.HandleFunc("/register", s.UserRegister).Methods(http.MethodPost)

	// admin routes
	// r.HandleFunc("/users", auth.WithJWTAuth(s.UsersGet, s.store)).Methods(http.MethodGet)
	// r.HandleFunc("/users/{userId}", auth.WithJWTAuth(s.UserGet, s.store)).Methods(http.MethodGet)
	r.HandleFunc("/users", s.UsersGet).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}", s.UserGet).Methods(http.MethodGet)
}

func (s *UserService) UserRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// _, err := s.store.GetUserByEmail(user.Email)
	_isvalid := s.store.RegistrationValidation(user)
	if !_isvalid {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	
	user_id, err := s.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Username: user.Username,
		Email: user.Email,
		Password: hashedPassword,
		Gender: user.Gender,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	// Token creation
	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, int(user_id))
	// err != nil {
	// 	utils.WriteError(w, http.StatusInternalServerError, err);
	// 	return
	// }
	
	// fmt.Print(token)

	s.store.UpdateUser(user_id, token);
	s.store.SendEmail("spamsama91@gmail.com");


	utils.WriteJSON(w, http.StatusCreated, user_id)
}

func (s *UserService) UserLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	_user, err := s.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	if !auth.ComparePasswords(_user.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(configs.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, _user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	is_valid := s.store.TokenValidation(token);
	if !is_valid {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, 1);


}

func (s *UserService) UsersGet(w http.ResponseWriter, r *http.Request) {
	users, err := s.store.GetAllUsers()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, utils.ErrorResponse{ Error: "Error getting users" })
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (s *UserService) UserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userId"]
	if !ok {
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("missing user Id"))
		return
	}

	userId, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("invalid user Id"))
		return
	}

	user, err := s.store.GetUserById(userId)
	if !ok {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
