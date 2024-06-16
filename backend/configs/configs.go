package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port       				string
	DBUser     				string
	DBPassword 				string
	DBAddress  				string
	DBName     				string
	JWTSecret  				string
	JWTExpiration			int64
	JWTExpirationInSeconds	int64
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:       			getEnv("PORT", "8000"),
		DBUser:     			getEnv("DB_USER", "root"),
		DBPassword: 			getEnv("DB_PASSWORD", ""),
		DBAddress:  			fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:     			getEnv("DB_NAME", "matcherx"),
		JWTSecret:  			getEnv("JWT_SECRET", "r4nd0mJW7$ecr3tkey"),
		JWTExpirationInSeconds:	getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 * 7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}