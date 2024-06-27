package store

import (
	"fmt"
	"log"
	"matchaVgo/internal/auth"

	// "matchaVgo/internal/auth"
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

func GetUserByEmail(db *sqlx.DB, email string) (*User, error) {

	var user User;
	err := db.Get(&user, "SELECT * FROM users WHERE email=$1", email);
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

func GetUserIDByEmail(db *sqlx.DB, email string) int64 {

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
        "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
        user.Username, user.Email, user.Password,
    ).Scan(&id)
    if err != nil {
        return -1, err
    }
    return id, nil
}

func UpdateUserToken(db *sqlx.DB, user *User, token string) (string, error) {
    
	// secret := []byte(os.Getenv("JWT_SECRET_TOKEN"));
	// token, err := auth.CreateJWT(int(user.ID));
    // if err != nil {
    //     log.Fatal(err);
    // }

	_, err := db.Exec("UPDATE users SET token = $1 WHERE id = $2", token, user.ID)
    if err != nil {
		return "", err
	}

	return token, nil;
}

func UpdateUser(db *sqlx.DB, user *User) (*User, error) {

	_, err := db.Exec("UPDATE users SET first_name = $1, last_name = $2, birthday = $3, gender = $4, preferences = $5, pics = $6, location = $7 WHERE id = $8", user.First_name, user.Last_name, user.Birthday, user.Gender, user.Preferences, user.Pics, user.Location, user.ID)
	if err != nil {
		return nil, err
	}

	updatedUser := &User{}
	err = db.Get(updatedUser, "SELECT * FROM users WHERE id = $1", user.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func UpdateUserPassword(db *sqlx.DB, user *ResetUserPassPayload) (*User, error) {

	hashedPassword, er := auth.HashPassword(user.Password)
	if er != nil {
        log.Fatalln(er);
	}

	_, err := db.Exec("UPDATE users SET password = $1 WHERE token = $2", hashedPassword, user.Token)
    if err != nil {
		return nil, err
	}

	updatedUser := &User{}
	err = db.Get(updatedUser, "SELECT * FROM users WHERE token = $1", user.Token)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func UpdateUserByToken(db *sqlx.DB, user *User) (*User, error) {

	_, err := db.Exec("UPDATE users SET first_name = $1, last_name = $2, birthday = $3, gender = $4, preferences = $5, pics = $6, location = $7 WHERE token = $8", user.First_name, user.Last_name, user.Birthday, user.Gender, user.Preferences, user.Pics, user.Location, user.Token)
	if err != nil {
		return nil, err
	}

	updatedUser := &User{}
	err = db.Get(updatedUser, "SELECT * FROM users WHERE token = $1", user.Token)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func SendEmail(user_email string) {

	email := "matcherx1337@gmail.com";
	email_pass := os.Getenv("EMAIL_PASS");
	front_url := os.Getenv("FRONT_URL");
	
	mail := gomail.NewMessage();
	mail.SetHeader("From", email);
	mail.SetHeader("To", user_email);
	mail.SetHeader("Subject", "MatcherX account verification");

	body := fmt.Sprintf(`<div><a href="%s"><b>Clicki 3la had lb3ar!</b></a> <br> <img src="%s" alt="img" /></div>`, front_url+"/proceed-signup", "https://media.makeameme.org/created/fact-no-verification.jpg");
	mail.SetBody("text/html", body);

	d := gomail.NewDialer("smtp.gmail.com", 587, email, email_pass);
	if err := d.DialAndSend(mail); err != nil {
		log.Print(err);
	}
	
}

func SendEmailPass(user_email string) (int, error) {

	email := "matcherx1337@gmail.com";
	email_pass := os.Getenv("EMAIL_PASS");
	front_url := os.Getenv("FRONT_URL");
	
	mail := gomail.NewMessage();
	mail.SetHeader("From", email);
	mail.SetHeader("To", user_email);
	mail.SetHeader("Subject", "MatcherX Reset Password");

	body := fmt.Sprintf(`<div><strong>nb9aw khdamin meak gha NTA?!</strong><a href="%s"><b>Clicki 3la had lb3ar!</b></a> <br> <img src="%s" alt="img" /></div>`, front_url+"/reset-pass", "https://i.pinimg.com/736x/f4/9c/c1/f49cc131e4c5c28c5697b15f67bb321f.jpg");
	mail.SetBody("text/html", body);

	d := gomail.NewDialer("smtp.gmail.com", 587, email, email_pass);
	if err := d.DialAndSend(mail); err != nil {
		log.Print(err);
		return -1, err;
	}
	
	return 1, nil;
}