package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type DB struct {
	*sqlx.DB
}

// type User struct {
// 	ID        int32  	`db:"id"`
// 	FirstName string 	`db:"first_name"`
// 	LastName  string 	`db:"last_name"`
// 	Email     string 	`db:"email"`
// 	Username  string 	`db:"username"`
// 	Password  string 	`db:"password"`
// 	Gender    bool		`db:"gender"`
// 	Token     string 	`db:"token"`
// }

// // NewDB creates a new database connection
// func NewDB(dataSourceName string) (*DB, error) {
// 	db, err := sqlx.Connect("postgres", dataSourceName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &DB{db}, nil
// }

func Connect() *sqlx.DB {

	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file");
	}

	db_name := os.Getenv("DB_NAME");
	db_user := os.Getenv("DB_USER");
	db_password := os.Getenv("DB_PASSWORD");

	s := "user=" + db_user + " password=" + db_password + " dbname=" + db_name + " sslmode=disable";
	db, err := sqlx.Connect("postgres", s);
	if err != nil {
		log.Fatalln(err);
	}

	return db;
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
