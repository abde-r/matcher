package main

import (
	"log"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User: Envs.DBUser,
		Password: Envs.DBPassword,
		addr: Envs.DBAddress,
		db: Envs.DBName,
		Net: "tcp",
		AllowNativePassword: true,
		ParseTime: true,
	}

	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)
	server := NewAPIServer(":1337", store)
	server.Run()
}
