package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port: getEnv("PORT"),
		DBUser: getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName: getEnv("DB_NAME"),
		JWTSecret: getEnv("JWT_SECRET"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return ""
}