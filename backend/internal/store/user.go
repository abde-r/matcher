package store

import (
	"fmt"
	"log"
	"matchaVgo/internal/auth"
	"os"

	"github.com/jmoiron/sqlx"
	"gopkg.in/gomail.v2"
)

func GetAllUsers(db *sqlx.DB) ([]User, error) {
    var users []User;
    err := db.Select(&users, "SELECT * FROM users");
    return users, err;
}

func GetUserById(db *sqlx.DB, id int64) (*User, error) {

	var user User;
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", id);
	if err != nil {
		return nil, err;
	}
	return &user, nil;
}

func GetUserByUsername(db *sqlx.DB, username string) int64 {

	var id int64
	err := db.QueryRow("SELECT id FROM users WHERE username=$1", username).Scan(&id)
	if err != nil {
		return -1
	}

	return id
}

func GetUserByEmail(db *sqlx.DB, email string) int64 {

	var id int64
	err := db.QueryRow("SELECT id FROM users WHERE email=$1", email).Scan(&id)
	if err != nil {
		return -1
	}

	return id
}

func CreateUser(db *sqlx.DB, user *User) (int32, error) {
    
	var id int32
    err := db.QueryRow(
        "INSERT INTO users (first_name, last_name, username, email, password, gender) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
        user.First_name, user.Last_name, user.Username, user.Email, user.Password, user.Gender,
    ).Scan(&id)
    if err != nil {
        return -1, err
    }
    return id, nil
}

func UpdateUserToken(db *sqlx.DB, user *User) (string, error) {
    
	secret := []byte(os.Getenv("JWT_SECRET_TOKEN"));
	token, err := auth.CreateJWT(secret, int(user.ID));
    if err != nil {
        log.Fatal(err);
    }

	_, err = db.Exec("UPDATE users SET token = $1 WHERE id = $2", token, user.ID)
    if err != nil {
		return "", err
	}

	return token, nil;
}

func UpdateUser(db *sqlx.DB, user *User) (*User, error) {

	id := "17"
	_, err := db.Exec("UPDATE users SET first_name = $1, last_name = $2, gender = $3 WHERE id = $4", user.First_name, user.Last_name, user.Gender, id/*user.ID*/)
	if err != nil {
		return nil, err
	}

	// Retrieve the updated user from the database
	updatedUser := &User{}
	err = db.Get(updatedUser, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func SendEmail(user_email string) {

	email := "matcherx1337@gmail.com";
	email_pass := os.Getenv("EMAIL_PASS");
	
	mail := gomail.NewMessage();
	mail.SetHeader("From", email);
	mail.SetHeader("To", user_email);
	mail.SetHeader("Subject", "MatcherX account verification");

	body := fmt.Sprintf(`<div><a href="%s"><b>Clicki 3la had lb3ar!</b></a> <br> <img src="%s" alt="img" /></div>`, "https://abder.vercel.app", "https://media.makeameme.org/created/fact-no-verification.jpg");
	mail.SetBody("text/html", body);

	d := gomail.NewDialer("smtp.gmail.com", 587, email, email_pass);
	if err := d.DialAndSend(mail); err != nil {
		log.Print(err);
	}
	
}