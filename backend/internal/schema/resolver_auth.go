package schema

import (
	"context"
	"errors"
	"log"
	"matchaVgo/internal/auth"
	"matchaVgo/internal/store"

)

// GraphQLUserRegistrationRequest represents the structure of a GraphQL query request
type GraphQLUserRegistrationRequest struct {
    Query     string                `json:"query" example:"mutation RegisterUser($input: RegisterUserInput!) { registerUser(input: $input) { username email password } }"`
    Variables store.RegisterUserPayload `json:"variables"`
}

// HTTPError represents the structure of an error response
type HTTPError struct {
    Code    int    `json:"code" example:"400"`
    Message string `json:"message" example:"Invalid input"`
}

// Matcher-doc
// @Summary User registeration
// @Description New user Registeration with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body GraphQLUserRegistrationRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/register [post]
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




// GraphQLUserLoginRequest represents the structure of a GraphQL query request
type GraphQLUserLoginRequest struct {
    Query     string                `json:"query" example:"mutation LoginUser($input: LoginUserInput!) { loginUser(input: $input) { username password } }"`
    Variables store.LoginUserPayload `json:"variables"`
}

// Matcher-doc
// @Summary User login
// @Description Existed user login with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body GraphQLUserLoginRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/login [post]
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


// GraphQLEmailVerificationRequest represents the structure of a GraphQL query request
type GraphQLEmailVerificationRequest struct {
    Query     string                `json:"query" example:"mutation SendEmailVerification($input: SendEmailVerificationPayload!) { sendEmailVerification(input: $input) { email } }"`
    Variables store.SendEmailVerificationPayload `json:"variables"`
}

// Matcher-doc
// @Summary Email verification
// @Description Send email verification to user by his email
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body GraphQLEmailVerificationRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/send-verification-email [post]
func (r *Resolver) SendEmailVerification(ctx context.Context, args struct{ Input store.SendEmailVerificationPayload }) (*UserResolver, error) {
	

	user, err := store.GetUserByEmail(db, args.Input.Email);
	if err != nil {
		return nil, errors.New("invalid email");
	}

    _, err = store.SendEmailPass(args.Input.Email);
	if err != nil {
        log.Fatal(err);
		return nil, err;
    }

	return &UserResolver{user: user}, nil;
}

// GraphQLPasswordResetRequest represents the structure of a GraphQL query request
type GraphQLPasswordResetRequest struct {
    Query     string                `json:"query" example:"mutation ResetUserPassword($input: ResetUserPassPayload!) { resetUserPassword(input: $input) { token } }"`
    Variables store.ResetUserPassPayload `json:"variables"`
}

// Matcher-doc
// @Summary Password reset
// @Description Reset password by user's token
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body GraphQLPasswordResetRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/reset-pass [post]
func (r *Resolver) ResetUserPassword(ctx context.Context, args struct{ Input store.ResetUserPassPayload }) (*UserResolver, error) {
	
	user, err := store.UpdateUserPassword(db, &args.Input);
	if err != nil {
        log.Fatal(err);
		return nil, err;
    }

	// return "1", nil;
	return &UserResolver{user: user}, nil;
} 