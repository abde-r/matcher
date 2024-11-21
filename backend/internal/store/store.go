package store

import (
	"log"

	// "time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

type User struct {
	ID          int32  `db:"id"`
	First_name  string `db:"first_name"`
	Last_name   string `db:"last_name"`
	Email       string `db:"email"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	Birthday    string `db:"birthday"`
	Gender      bool   `db:"gender"`
	Preferences string `db:"preferences"`
	Pics        string `db:"pics"`
	Location    string `db:"location"`
	Token       string `db:"token"`
}

// RegisterUserPayload represents the payload for user registration
type RegisterUserPayload struct {
	Username string `json:"username" validate:"required" example:"ms3oud"`
	Email    string `json:"email" validate:"required,email" example:"ms3oud@example.test"`
	Password string `json:"password" validate:"required,min=6,max=20" example:"securePASSWORD123"`
	Token    string `json:"token"`
}

type ProceedRegistrationUserPayload struct {
	ID          int32  `db:"id"`
	Username    string `db:"username"`
	Email       string `db:"email"`
	First_name  string `db:"first_name" validate:"required"`
	Last_name   string `db:"last_name" validate:"required"`
	Birthday    string `db:"birthday" validate:"required"`
	Gender      bool   `db:"gender" validate:"required"`
	Preferences string `db:"preferences" validate:"required"`
	Pics        string `db:"pics" validate:"required"`
	Location    string `db:"location" validate:"required"`
	Token       string `db:"token" validate:"required"`
}

type UpdateUserInfoPayload struct {
	ID          int32  `db:"id"`
	Username    string `db:"username"`
	Email       string `db:"email"`
	First_name  string `db:"first_name" validate:"required"`
	Last_name   string `db:"last_name" validate:"required"`
	Birthday    string `db:"birthday" validate:"required"`
	Preferences string `db:"preferences" validate:"required"`
	Pics        string `db:"pics" validate:"required"`
	Location    string `db:"location" validate:"required"`
	Token       string `db:"token" validate:"required"`
}

type LoginUserPayload struct {
	Username string `json:"username" validate:"required" example:"ms3oud"`
	Password string `json:"password" validate:"required" example:"securePASSWORD123"`
}

type SendEmailVerificationPayload struct {
	Email string `json:"email" validate:"required,email" example:"ms3oud@example.test"`
}

type ResetUserPassPayload struct {
	Password string `json:"password" validate:"required"`
	Token    string `json:"token" validate:"required"`
}

type MessagePayload struct {
	ID         string `json:"id" validate:"required"`
	SenderID   string `json:"sender_id" validate:"required"`
	ReceiverID string `json:"receiver_id" validate:"required"`
	Content    string `json:"content" validate:"required"`
	Timestamp  string `json:"timestamp" validate:"required"`
}

type TypingEvent struct {
	UserID string `json:"user_id" validate:"required"`
	Typing bool   `json:"typing" validate:"required"`
}

func GetDB() *sqlx.DB {
	if db != nil {
		log.Fatalln("Database connection is not initialized")
	}

	return db
}
