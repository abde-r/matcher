package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB;

type User struct {
    ID        	int32  	`db:"id"`
    First_name 	string 	`db:"first_name"`
    Last_name  	string 	`db:"last_name"`
    Email     	string 	`db:"email"`
    Username  	string 	`db:"username"`
    Password  	string 	`db:"password"`
    Gender    	bool 	`db:"gender"`
	Token		string	`db:"token"`
}

type RegisterUserPayload struct {
	First_name	string  	`json:"first_name"`
	Last_name	string  	`json:"last_name"`
	Username	string  	`json:"username" validate:"required"`
	Email		string  	`json:"email" validate:"required,email"`
	Password	string  	`json:"password" validate:"required,min=6,max=20"`
	Gender		bool  		`json:"gender"`
	Token		string  	`json:"token"`
}

type CompleteRegistrationUserPayload struct {
	First_name	string  	`json:"first_name" validate:"required"`
	Last_name	string  	`json:"last_name" validate:"required"`
	Username	string  	`json:"username"`
	Email		string  	`json:"email"`
	Password	string  	`json:"password"`
	Gender		bool  		`json:"gender" validate:"required"`
	Token		string  	`json:"token"`
}

type LoginUserPayload struct {
	Username	string	`json:"username" validate:"required"`
	Password	string	`json:"password" validate:"required"`
}

func GetDB() *sqlx.DB {
	if db != nil {
		log.Fatalln("Database connection is not initialized");
	}

	return db;
}