package db

import (
	"log"
	"os"
    "path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type DB struct {
	*sqlx.DB
}

func Connect() *sqlx.DB {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	db_name := os.Getenv("POSTGRES_DB")
	db_user := os.Getenv("POSTGRES_USER")
	db_password := os.Getenv("POSTGRES_PASSWORD")

	connString := "host=" + db_host + " port=" + db_port + " user=" + db_user +
        		" password=" + db_password + " dbname=" + db_name + " sslmode=disable"

	println(connString)
    db, err := sqlx.Connect(db_user, connString)
    if err != nil {
        log.Fatalln("Database connection error:", err)
    }

	return db
}

// // CreateUser inserts a new user into the database and returns the user's ID
// func (db *DB) CreateUser(user *User) (int32, error) {
// 	var id int32
// 	query := `INSERT INTO users (first_name, last_name, email, username, password, gender, token)
// 			  VALUES (:first_name, :last_name, :email, :username, :password, :gender, :token) RETURNING id`
// 	stmt, err := db.NamedQuery(query, user)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()
// 	if stmt.Next() {
// 		if err := stmt.Scan(&id); err != nil {
// 			return 0, err
// 		}
// 	}
// 	return id, nil
// }
