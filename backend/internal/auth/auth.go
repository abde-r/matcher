package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWT(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(3600 * 24 * 7);

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(int(userId)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

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