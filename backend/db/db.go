package db

import (
	"database/sql"
	"log"
	
	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func (s *MySQLStorage) Init() (*sql.DB, error) {

	// Initialize the tables
	return s.db, nil
}
