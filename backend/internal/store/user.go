package store

import (
    "github.com/jmoiron/sqlx"
)

type User struct {
    ID        	int32  	`db:"id"`
    FirstName 	string 	`db:"first_name"`
    LastName  	string 	`db:"last_name"`
    Email     	string 	`db:"email"`
    Username  	string 	`db:"username"`
    Password  	string 	`db:"password"`
    Gender    	string 	`db:"gender"`
	Token		string	`db:"token"`
}

func GetAllUsers(db *sqlx.DB) ([]User, error) {
    var users []User
    err := db.Select(&users, "SELECT * FROM users")
    return users, err
}

func CreateUser(db *sqlx.DB, user *User) (int32, error) {
    var id int32
    err := db.QueryRow(
        "INSERT INTO users (first_name, last_name, username, email, password, gender) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
        user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Gender,
    ).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}
