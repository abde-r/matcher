package types

import "time"

type User struct {
	Id        	int  		`json:"id"`
	FirstName	string  	`json:"firstName"`
	LastName	string  	`json:"lastName"`
	Username	string  	`json:"username"`
	Email		string  	`json:"email"`
	Password	string  	`json:"-"`
	Gender		bool  		`json:"gender"`
}

type RegisterUserPayload struct {
	FirstName	string  	`json:"firstName" validate:"required"`
	LastName	string  	`json:"lastName" validate:"required"`
	Username	string  	`json:"username" validate:"required"`
	Email		string  	`json:"email" validate:"required,email"`
	Password	string  	`json:"password" validate:"required,min=3,max=10"`
	Gender		bool  		`json:"gender" validate:"required"`
}

type Post struct {
	Id        	int  	`json:"id"`
	Title   	string 	`json:"title"`
	CreatedAt	time.Time `json:"createdAt"`
}

type UserStore interface {
	CreateUser(User) error
	GetAllUsers() ([]User, error)
	GetUserById(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(email string) (*User, error)
	RegistrationValidation(RegisterUserPayload) bool
}