package user

import (
	// "bytes"
	"database/sql"
	"fmt"
	"log"
	"matchaVgo/types"
	// "os"

	// "matchaVgo/services/user/template"
	// "html/template"
	// "path/filepath"
	// "runtime"

	"gopkg.in/gomail.v2"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllUsers() ([]types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []types.User
    for rows.Next() {
		var user types.User
        if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Gender, &user.Token); err != nil {
			log.Fatal(err)
			return users, err
        }
        users = append(users, user)
    }
    if err = rows.Err(); err != nil {
        return users, err
    }
    return users, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Store) TokenValidation(token string) bool {
	rows, err := s.db.Query("SELECT * FROM user WHERE token = ?", token)
	if err != nil {
		return false
	}

	defer rows.Close()

	for rows.Next() {
		user := new(types.User)
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Gender, &user.Token)
		if err != nil {
			// log.Printf("Error scanning row into user: %v", err)
			return false
		}
		// If we successfully find a user, return true
		return true
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		// log.Printf("Error encountered during row iteration: %v", err)
		return false
	}

	// If no user is found, return false
	return false
}

func (s *Store) GetUserById(id int) (*types.User, error) {

	rows, err := s.db.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil

}

func (s *Store) CreateUser(user types.User) (int64, error) {

	_, err := s.db.Exec("CREATE TABLE IF NOT EXISTS user (id INT NOT NULL AUTO_INCREMENT, firstName TEXT, lastName TEXT, username TEXT, email TEXT, password TEXT, gender TEXT, token TEXT, PRIMARY KEY (id))");
	if err != nil {
		return -1, fmt.Errorf("error creating table 'user' : %v", err);
	}

	res, err := s.db.Exec("INSERT INTO user (firstName, lastName, username, email, password, gender, token) VALUES (?,?,?,?,?,?,?)", user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Gender, "token")
	if err != nil {
		return -1, fmt.Errorf("error inserting user : %v", err);
	}

	user_id, err := res.LastInsertId();
	if err != nil {
		return -1, fmt.Errorf("error getting last user id : %v", err);
	}

	return user_id, nil
}

func (s *Store) UpdateUser(userId int64, token string) (string, error) {

	_, err := s.db.Exec("UPDATE user SET token = ? WHERE id = ?", token, userId);
	if err != nil {
		return token, fmt.Errorf("error updating table 'user' : %v", err);
	}
	return token, err;
}

var token = "vrck zbeh opcy orpm";

func (s *Store) SendEmail(user_email string) {

	// email, exists := os.LookupEnv("EMAIL_PASS");
	// if !exists {
	// 	fmt.Print("env variable not found");
	// }
	// pass, exists := os.LookupEnv("EMAIL_PASS");
	// if !exists {
	// 	fmt.Print("env variable not found");
	// }
	email := "matcherx1337@gmail.com";
	
	mail := gomail.NewMessage();
	mail.SetHeader("From", email);
	mail.SetHeader("To", user_email);
	mail.SetHeader("Subject", "MatcherX account verification");

	body := fmt.Sprintf(`<a href="%s"><b>Clicki 3la had lb3ar!</b></a> <br> <img src="%s" alt="img" />`, "https://abder.vercel.app", "https://media.makeameme.org/created/fact-no-verification.jpg");
	mail.SetBody("text/html", body);

	d := gomail.NewDialer("smtp.gmail.com", 587, email, token);
	if err := d.DialAndSend(mail); err != nil {
		log.Print(err);
	}
	
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Gender, &user.Token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) RegistrationValidation(user types.RegisterUserPayload) bool {
	
	_, err := s.GetUserByEmail(user.Email)
	if err == nil {
		return false
	}

	_, err = s.GetUserByUsername(user.Username)
	return err != nil
	
	// return true
}