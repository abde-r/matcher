package schema

import (
	"context"

	"matchaVgo/internal/store"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

// UserResolver struct
type UserResolver struct {
	user *store.User
}

// ErrorResponse
type ErrorResponse struct {
	Message string `json:"message"`
}

// Field resolvers for User
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.user.ID)))
}

func (r *UserResolver) FirstName() string {
	return r.user.First_name
}

func (r *UserResolver) LastName() string {
	return r.user.Last_name
}

func (r *UserResolver) Email() string {
	return r.user.Email
}

func (r *UserResolver) Username() string {
	return r.user.Username
}

func (r *UserResolver) Password() string {
	return r.user.Password
}

func (r *UserResolver) Gender() bool {
	return r.user.Gender
}

func (r *UserResolver) Birthday() string {
	return r.user.Birthday
}

func (r *UserResolver) Preferences() string {
	return r.user.Preferences
}

func (r *UserResolver) Pics() string {
	return r.user.Pics
}

func (r *UserResolver) Token() string {
	return r.user.Token
}

func (r *UserResolver) Location() string {
	return r.user.Location
}

// GraphQLUsersRequest represents the structure of a GraphQL query request
type GraphQLUsersRequest struct {
	Query     string   `json:"query" example:"mutation Users($input: User!) { user(input: $input) { } }"`
	Variables struct{} `json:"variables"`
}

// Matcher-doc
// @Summary Users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Param input body GraphQLUsersRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/ [post]
func (r *Resolver) Users(ctx context.Context) ([]*UserResolver, error) {
	var users []store.User
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	userResolvers := make([]*UserResolver, len(users))
	for i := range users {
		userResolvers[i] = &UserResolver{user: &users[i]}
	}
	return userResolvers, nil
}

// GraphQLUserByIDRequest represents the structure of a GraphQL query request
type GraphQLUserByIDRequest struct {
	Query     string            `json:"query" example:"mutation User($input: UserByIDVariables!) { user(input: $input) { id } }"`
	Variables UserByIDVariables `json:"variables"`
}

// UserByIDVariables represents the variables passed to the GraphQL request
type UserByIDVariables struct {
	ID string `json:"ID"`
}

// Matcher-doc
// @Summary User by ID
// @Description Get user by ID
// @Tags User
// @Accept json
// @Produce json
// @Param input body GraphQLUserByIDRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/id [post]
func (r *Resolver) User(ctx context.Context, args struct{ ID int32 }) (*UserResolver, error) {
	var user store.User
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", args.ID)
	if err != nil {
		return nil, err
	}
	return &UserResolver{user: &user}, nil
}

// GraphQLUserByTokenRequest represents the structure of a GraphQL query request
type GraphQLUserByTokenRequest struct {
	Query     string               `json:"query" example:"mutation UserByToken($input: UserByTokenVariables!) { userByToken(input: $input) { token } }"`
	Variables UserByTokenVariables `json:"variables"`
}

// UserByTokenVariables represents the variables passed to the GraphQL request
type UserByTokenVariables struct {
	Token string `json:"token"`
}

// Matcher-doc
// @Summary User by Token
// @Description Get user by token
// @Tags User
// @Accept json
// @Produce json
// @Param input body GraphQLUserByTokenRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/token [post]
func (r *Resolver) UserByToken(ctx context.Context, token string) (*UserResolver, error) {
	var user store.User
	err := db.Get(&user, "SELECT * FROM users WHERE token=$1", token)
	if err != nil {
		return nil, err
	}
	return &UserResolver{user: &user}, nil
}

// GraphQLProceedRegistrationRequest represents the structure of a GraphQL query request
type GraphQLProceedRegistrationRequest struct {
	Query     string                               `json:"query" example:"mutation ProceedRegistrationUser($input: ProceedRegistrationUserPayload!) { proceedRegistrationUser(input: $input) { first_name last_name birthday gender preferences pics location token} }"`
	Variables store.ProceedRegistrationUserPayload `json:"variables"`
}

// Matcher-doc
// @Summary Proceed registration
// @Description Proceed registration of user
// @Tags User
// @Accept json
// @Produce json
// @Param input body GraphQLProceedRegistrationRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/proceed-registration [post]
func (r *Resolver) ProceedRegistrationUser(ctx context.Context, args struct {
	Input store.ProceedRegistrationUserPayload
}) (*UserResolver, error) {

	user := store.User{
		First_name:  args.Input.First_name,
		Last_name:   args.Input.Last_name,
		Birthday:    args.Input.Birthday,
		Gender:      args.Input.Gender,
		Preferences: args.Input.Preferences,
		Pics:        args.Input.Pics,
		Location:    args.Input.Location,
		Token:       args.Input.Token,
	}

	_user, err := store.UpdateUserByToken(db, &user)
	if err != nil {
		return nil, err
	}

	return &UserResolver{user: _user}, nil
}

// GraphQLUpdateUserRequest represents the structure of a GraphQL query request
type GraphQLUpdateUserRequest struct {
	Query     string                      `json:"query" example:"mutation UpdateUserInfo($input: UpdateUserInfoPayload!) { updateUserInfo(input: $input) { first_name last_name birthday gender preferences pics location token} }"`
	Variables store.UpdateUserInfoPayload `json:"variables"`
}

// Matcher-doc
// @Summary Proceed registration
// @Description Proceed registration of user
// @Tags User
// @Accept json
// @Produce json
// @Param input body GraphQLUpdateUserRequest true "GraphQL Mutation Payload"
// @Success 200 {object} store.User
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/update-info [post]
func (r *Resolver) UpdateUserInfo(ctx context.Context, args struct{ Input store.UpdateUserInfoPayload }) (*UserResolver, error) {

	user := store.User{
		First_name:  args.Input.First_name,
		Last_name:   args.Input.Last_name,
		Birthday:    args.Input.Birthday,
		Preferences: args.Input.Preferences,
		Pics:        args.Input.Pics,
		Location:    args.Input.Location,
		Token:       args.Input.Token,
	}

	_user, err := store.UpdateUserByToken(db, &user)
	if err != nil {
		return nil, err
	}

	return &UserResolver{user: _user}, nil
}
