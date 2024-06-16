package auth

import (
	"context"
	"fmt"
	"log"
	"matchaVgo/configs"
	"matchaVgo/types"
	"matchaVgo/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type contextKey string
const UserKey contextKey = "userId"

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		token, err := validateJWT(tokenString)

		if err != nil {
			log.Println("Failed to authenticate Token")
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("Invalide Token")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)
		userId, err := strconv.Atoi(str)
		if err != nil {
			log.Println("conversion faild!", err)
			permissionDenied(w)
			return
		}

		_user, err := store.GetUserById(userId)
		if err != nil {
			log.Println("User not found!")
			permissionDenied(w)
			return
		}

		// Add user to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, _user.Id)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteJSON(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	} else if tokenQuery != "" {
		return tokenQuery
	} else {
		return ""
	}
}

func validateJWT(t string) (*jwt.Token, error) {
	secret := configs.Envs.JWTSecret

	return jwt.Parse(t, func(t *jwt.Token) (interface {}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func CreateJWT(secret []byte, userId int) (string, error) {

	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(int(userId)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})
	
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}