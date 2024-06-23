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
    newUser.Token, err = store.UpdateUserToken(db, &newUser);
	if err != nil {
        log.Fatal(err);
    }

    store.SendEmail("spamsama91@gmail.com");
	

	return &UserResolver{user: &newUser}, nil
}

func (r *Resolver) LoginUser(ctx context.Context, args struct{ Input store.LoginUserPayload }) (*UserResolver, error) {
	
	user, err := store.LoginValidation(db, &args.Input);
	if err != nil {
		return nil, err;
	}

	// JWT or something

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