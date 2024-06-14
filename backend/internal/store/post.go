package store

import (
    "github.com/jmoiron/sqlx"
)

type Post struct {
	ID      int32  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	UserID  int    `db:"user_id"`
}

func GetAllPosts(db *sqlx.DB) ([]Post, error) {
    var posts []Post
    err := db.Select(&posts, "SELECT * FROM posts")
    return posts, err
}

func CreatePost(db *sqlx.DB, post *Post) (int32, error) {
    var id int32
    err := db.QueryRow(
        "INSERT INTO posts (title, content, user_id) VALUES ($1, $2, $3) RETURNING id",
        post.Title, post.Content, post.UserID,
    ).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}
