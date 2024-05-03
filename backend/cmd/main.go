package main

import (
	"database/sql"
	"fmt"
	"log"
	"matchaVgo/cmd/api"
	"matchaVgo/configs"
	"matchaVgo/db"

	"github.com/go-sql-driver/mysql"
)

func main() {

	cfg := mysql.Config{
		User: configs.Envs.DBUser,
		Passwd: configs.Envs.DBPassword,
		Addr: configs.Envs.DBAddress,
		DBName: configs.Envs.DBName,
		// Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}

	db, err := db.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	// store := NewStore(db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Database Successfully connected!")
}