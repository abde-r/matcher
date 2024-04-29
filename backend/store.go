package main

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

type Store interface {

	// Users
	CreateUser(user *User) (*User, error)
	GetUserById(id string) (*User, error)

	// Posts
	GetAllPosts() ([]Post, error)
	GetPostById(id string) (*Post, error)
	CreatePost(post *Post) error
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// *** User Services
func (s *Storage) GetUserById(id string) (*User, error) {
	
	var user User
	err := s.db.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(&user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Password, &user.Gender)
	return &user, err

}

func (s *Storage) CreateUser(user *User) (*User, error) {
	
	rows, err := s.db.Exec("INSERT INTO user (firstName, lastName, username, email, password, gender) VALUES (?,?,?,?,?,?)", user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Gender)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.Id = id
	return user, nil
}

// *** Posts Services
func (s *Storage) GetAllPosts() ([]Post, error) {
	
	rows, err := s.db.Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.Id, &p.Title)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return posts, nil

}

func (s *Storage) GetPostById(id string) (*Post, error) {
	
	var post Post
	err := s.db.QueryRow("SELECT * FROM post WHERE id = ?", id).Scan(&post.Id, &post.Title)
	return &post, err

}

func (s *Storage) CreatePost(post *Post) error {

	_, err := s.db.Exec("INSERT INTO post (title) VALUES (?)", post.Title)
	return err
}