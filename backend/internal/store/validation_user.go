package store

import (
	"errors"
	"matchaVgo/internal/auth"

	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
)

var Validate = validator.New()

func RegistrationValidation(db *sqlx.DB, user *RegisterUserPayload) (bool, error) {

	if err := Validate.Struct(user); err != nil {
		return false, err;
	}

	if GetUserByUsername(db, user.Username) != -1 {
		return false, errors.New("user already exists");
	}

	if GetUserByEmail(db, user.Email) != -1 {
		return false, errors.New("user already exists");
	}

	return true, nil;

}

func CompleteRegistrationValidation(db *sqlx.DB, user *CompleteRegistrationUserPayload) (bool, error) {

	if err := Validate.Struct(user); err != nil {
		return false, err;
	}

	if GetUserByUsername(db, user.Username) != -1 {
		return false, errors.New("user already exists");
	}

	if GetUserByEmail(db, user.Email) != -1 {
		return false, errors.New("user already exists");
	}

	return true, nil;

}

func LoginValidation(db *sqlx.DB, user *LoginUserPayload) (*User, error) {

	if err := Validate.Struct(user); err != nil {
		return nil, err;
	}

	id := GetUserByUsername(db, user.Username);
	if id == -1 {
		return nil, errors.New("invalid email or password");
	}
	
	_user, err := GetUserById(db, id);
	if err != nil {
		return nil, errors.New("error getting user by id");
	}
	
	if !auth.ComparePasswords(_user.Password, []byte(user.Password)) {
		return nil, errors.New("invalid email or password");
	}

	return _user, nil;

}