package schema

import (
	// "context"
	// "log"
	// "matchaVgo/internal/db"
	// "matchaVgo/services/auth"
	// "matchaVgo/utils"
	// "os"
	// "strconv"

	// "github.com/graph-gophers/graphql-go"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetDB(database *sqlx.DB) {
	db = database
}



// // User struct
// type User struct {
// 	ID        int32  `db:"id"`
// 	FirstName string `db:"first_name"`
// 	LastName  string `db:"last_name"`
// 	Email     string `db:"email"`
// 	Username  string `db:"username"`
// 	Password  string `db:"password"`
// 	Gender    string `db:"gender"`
// 	Token     string `db:"token"`
// }

// Post struct
// type Post struct {
// 	ID      int32  `db:"id"`
// 	Title   string `db:"title"`
// 	Content string `db:"content"`
// 	UserID  int    `db:"user_id"`
// }

// CreateUserInput struct
type CreateUserInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Token     string `json:"token"`
}

// Query resolver methods
