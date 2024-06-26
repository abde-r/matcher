package schema

import (
	"context"
	"log"
	"matchaVgo/internal/auth"
	"matchaVgo/internal/store"
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

	return &UserResolver{user: user}, nil;
}