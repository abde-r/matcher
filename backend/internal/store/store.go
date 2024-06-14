package store

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB;

func GetDB() *sqlx.DB {
	if db != nil {
		log.Fatalln("Database connection is not initialized");
	}

	return db;
}