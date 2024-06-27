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
		Email:     args.Input.Email,
		Username:  args.Input.Username,
		Password:  hashedPassword,
	}

	id, err := store.CreateUser(db, &newUser);
	if err != nil {
		return nil, err
	}
	
	newUser.ID = id
	token, err := auth.SetCookiza(ctx, int(newUser.ID));
	if err != nil {
        log.Fatal(err);
    }

	store.UpdateUserToken(db, &newUser, token);

    store.SendEmail("spamsama91@gmail.com");

	return &UserResolver{user: &newUser}, nil
}


func (r *Resolver) LoginUser(ctx context.Context, args struct{ Input store.LoginUserPayload }) (*UserResolver, error) {
	
	user, err := store.LoginValidation(db, &args.Input);
	if err != nil {
		return nil, err;
	}

	token, err := auth.SetCookiza(ctx, int(user.ID));
	if err != nil {
        log.Fatal(err);
    }

	store.UpdateUserToken(db, user, token);

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
}