package auth

import (
	// "os"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWT(userId int) (string, error) {

	// expiration := time.Second * time.Duration(3600*24*7)
	secret := os.Getenv("JWT_SECRET_TOKEN");
	jwt_secret := []byte(secret);
	expiration := time.Minute * 3;


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(int(userId)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(jwt_secret);
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// func generateJWT(user_id int) (string, error) {
// 	// Create the Claims
// 	jwt_secret := []byte(os.Getenv("JWT_SECRET_TOKEN"))
// 	claims := &jwt.StandardClaims{
// 		Subject:   strconv.Itoa(user_id),
// 		ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
// 	}

// 	// Create the token using your secret key
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwt_secret)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func ComparePasswords(hashed string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plain)
	return err == nil
}
