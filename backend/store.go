package main

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

type Store interface {

	// Users
	CreateUser() error

	// Posts
	GetAllPosts() ([]Post, error)
	GetPost(id string) (*Post, error)
	CreatePost(p *Post) error
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// *** User


// *** Posts
func (s *Storage) GetAllPosts() ([]Post, error) {
	
	rows, err := s.db.Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.Title)
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

func (s *Storage) GetPost(id string) (*Post, error) {
	
	var p Post
	err := s.db.QueryRow("SELECT * FROM post WHERE id = ?", id).Scan(&p.ID, &p.Title)
	return &p, err

}

func (s *Storage) CreatePost(p *Post) error {

	_, err := s.db.Exec("INSERT INTO post (title) VALUES (?)", p.Title)
	return err
}