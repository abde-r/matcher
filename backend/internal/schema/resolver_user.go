package schema

import (
	"context"
	"fmt"
	// "time"

	// "fmt"
	// "io/ioutil"
	// "matchaVgo/internal/schema"
	"matchaVgo/internal/store"
	// "net/http"
	"strconv"
	// "strings"

	"github.com/graph-gophers/graphql-go"
)

// UserResolver struct
type UserResolver struct {
	user *store.User
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
	return r.user.Birthday;
}

func (r *UserResolver) Preferences() string {
	return r.user.Preferences;
}

func (r *UserResolver) Pics() string {
	return r.user.Pics;
}

func (r *UserResolver) Token() string {
	return r.user.Token;
}

func (r *UserResolver) Location() string {
	return r.user.Location;
}

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

func (r *Resolver) User(ctx context.Context, args struct{ ID int32 }) (*UserResolver, error) {
	var user store.User
	err := db.Get(&user, "SELECT * FROM users WHERE id=$1", args.ID)
	if err != nil {
		return nil, err
	}
	return &UserResolver{user: &user}, nil
}

func (r *Resolver) UserByToken(ctx context.Context, args struct{ Token string }) (*UserResolver, error) {
	var user store.User
	err := db.Get(&user, "SELECT * FROM users WHERE token=$1", args.Token)
	if err != nil {
		return nil, err
	}
	return &UserResolver{user: &user}, nil
}

func (r *Resolver) ProceedRegistrationUser(ctx context.Context, args struct{ Input store.ProceedRegistrationUserPayload }) (*UserResolver, error) {
	
	user := store.User{
		ID:		args.Input.ID,
		First_name: args.Input.First_name,
		Last_name:  args.Input.Last_name,
		Birthday:    args.Input.Birthday,
		Gender:    args.Input.Gender,
		Preferences:    args.Input.Preferences,
		Pics: args.Input.Pics,
		Location: args.Input.Location,
	}

	fmt.Println("hello", user)
	_user, err := store.UpdateUser(db, &user);	
	if err != nil {
		return nil, err
	}
	// id, err := strconv.Atoi(string(args.ID))
	// if err != nil {
	// 	return nil, err
	// }

	// user, err := store.GetUserById(db, int64(id));
	// if err != nil {
	// 	return nil, err
	// }
	
	// _user, err := store.GetUserById(db, 17);
	// if err != nil {
	// 	return nil, err
	// }

	// Perform any additional registration completion steps here

	return &UserResolver{user: _user}, nil
}