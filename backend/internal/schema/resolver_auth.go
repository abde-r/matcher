package schema

import (
	"context"
	"fmt"
	// "os"

	// "os"

	// "fmt"
	"log"
	"matchaVgo/internal/auth"
	"matchaVgo/internal/store"
	"net/http"
	"time"
	// "github.com/graph-gophers/graphql-go"
	// "github.com/99designs/gqlgen/graphql"
)

func (r *Resolver) RegisterUser(ctx context.Context, args struct{ Input store.RegisterUserPayload }) (*UserResolver, error) {

	is_valid, err := store.RegistrationValidation(db, &args.Input);
	if !is_valid || err != nil {
		return nil, err;
	}

    hashedPassword, er := auth.HashPassword(args.Input.Password)
	if er != nil {
        log.Fatalln(er);
	}
    
    newUser := store.User{
		// First_name: args.Input.First_name,
		// Last_name:  args.Input.Last_name,
		Email:     args.Input.Email,
		Username:  args.Input.Username,
		Password:  hashedPassword,
		// Gender:    args.Input.Gender,
	}

	id, err := store.CreateUser(db, &newUser);
	if err != nil {
		return nil, err
	}
	
	newUser.ID = id
    newUser.Token, err = auth.CreateJWT(int(newUser.ID));// store.UpdateUserToken(db, &newUser);
	if err != nil {
        log.Fatal(err);
    }

    store.SendEmail("spamsama91@gmail.com");
	

	return &UserResolver{user: &newUser}, nil
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

func (r *Resolver) LoginUser(ctx context.Context, args struct{ Input store.LoginUserPayload }) (*UserResolver, error) {
	
	user, err := store.LoginValidation(db, &args.Input);
	if err != nil {
		return nil, err;
	}

	token, err := auth.CreateJWT(int(user.ID))
	if err != nil {
		return nil, err
	}

	// Extract the HTTP response writer from the context
	if httpResponseWriter, ok := getResponseWriter(ctx); ok {
		// Set the cookie
		http.SetCookie(httpResponseWriter, &http.Cookie{
			Name:     "matcher-token",
			Value:    token,
			Path:     "/",
			Expires:  time.Now().Add(time.Minute * 3), // Cookie expires in 1 minute
			HttpOnly: false,
			Secure:   false, // Set to true if using HTTPS
			SameSite: http.SameSiteStrictMode,
		})
		store.UpdateUserToken(db, user, token);
	} else {
		return nil, err;
	}
		
		fmt.Println("Hola", token);
	return &UserResolver{user: user}, nil;

	// var user types.LoginUserPayload
	// if err := utils.ParseJSON(r, &user); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

	// if err := utils.Validate.Struct(user); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	// 	return
	// }

	// _user, err := s.store.GetUserByEmail(user.Email)
	// if err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
	// 	return
	// }

	// if !auth.ComparePasswords(_user.Password, []byte(user.Password)) {
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
	// 	return
	// }

	// secret := []byte(configs.Envs.JWTSecret)
	// token, err := auth.CreateJWT(secret, _user.Id)
	// if err != nil {
	// 	utils.WriteError(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// is_valid := s.store.TokenValidation(token);
	// if !is_valid {
	// 	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
	// 	return
	// }

	// utils.WriteJSON(w, http.StatusOK, 1);
}