package user

import (
	"database/sql"
	"fmt"
	"matchaVgo/types"
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
        if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Gender); err != nil {
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

func (s *Store) CreateUser(user types.User) error {

	_, err := s.db.Exec("INSERT INTO user (firstName, lastName, username, email, password, gender) VALUES (?,?,?,?,?,?)", user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Gender)
	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Gender)
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