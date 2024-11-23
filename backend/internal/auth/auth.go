package auth

import (
	// "os"
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func CreateJWT(userId int) (string, error) {

	// expiration := time.Second * time.Duration(3600*24*7)
	secret := os.Getenv("JWT_SECRET_TOKEN")
	jwt_secret := []byte(secret)
	expiration := time.Minute * 15

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(int(userId)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(jwt_secret)
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

type contextKey string

const responseWriterKey contextKey = "responseWriter"

func WithResponseWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), responseWriterKey, w)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getResponseWriter(ctx context.Context) (http.ResponseWriter, bool) {
	w, ok := ctx.Value(responseWriterKey).(http.ResponseWriter)
	return w, ok
}

func SetCookiza(ctx context.Context, user_id int) (string, error) {

	token, err := CreateJWT(int(user_id))
	if err != nil {
		return "", err
	}

	// Extract the HTTP response writer from the context
	if httpResponseWriter, ok := getResponseWriter(ctx); ok {
		http.SetCookie(httpResponseWriter, &http.Cookie{
			Name:     "matcher-token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(time.Minute * 15), // Cookie expires in 1 minute
			HttpOnly: false,                            // TO UPDATE LATER, must be true
			Secure:   false,                            // Set to true if using HTTPS
			SameSite: http.SameSiteStrictMode,
		})
	} else {
		return "", err
	}
	return token, nil
}

func EncryptEncryptedToken(encryptedtoken string) (string, error) {
	// expiration := time.Second * time.Duration(3600*24*7)
	secret := os.Getenv("JWT_SECRET_TOKEN")
	jwt_secret := []byte(secret)
	expiration := time.Minute * 15

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    encryptedtoken,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(jwt_secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DycreptEncryptedToken(encryptedtoken string) (string, error) {

	return "", nil
}
