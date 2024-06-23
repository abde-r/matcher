package store

import (
	"log"

	// "time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB;


type User struct {
    ID        	int32  		`db:"id"`
    First_name 	string 		`db:"first_name"`
    Last_name  	string 		`db:"last_name"`
    Email     	string 		`db:"email"`
    Username  	string 		`db:"username"`
    Password  	string 		`db:"password"`
	Birthday	string		`db:"birthday"`
    Gender    	bool		`db:"gender"`
	Preferences	string		`db:"preferences"`
	Pics		string		`db:"pics"`
	Location	string		`db:"location"`
	Token		string		`db:"token"`
}

type RegisterUserPayload struct {
	Username	string  	`json:"username" validate:"required"`
	Email		string  	`json:"email" validate:"required,email"`
	Password	string  	`json:"password" validate:"required,min=6,max=20"`
	Token		string  	`json:"token"`
}

type ProceedRegistrationUserPayload struct {
    ID        	int32  		`db:"id"`
	Username	string  	`db:"username"`
	Email		string  	`db:"email"`
	First_name	string  	`db:"first_name" validate:"required"`
	Last_name	string  	`db:"last_name" validate:"required"`
	Birthday	string  	`db:"birthday" validate:"required"`
	Gender		bool  		`db:"gender" validate:"required"`
	Preferences	string  	`db:"preferences" validate:"required"`
	Pics		string		`db:"pics" validate:"required"`
	Location	string		`db:"location" validate:"required"`
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