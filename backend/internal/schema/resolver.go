package schema

import (
	"context"
	"log"
	"matchaVgo/internal/auth"
	"matchaVgo/internal/store"
	"matchaVgo/utils"
	"os"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

// Resolver struct
type Resolver struct{}

// UserResolver struct
type UserResolver struct {
	user *store.User
}

// PostResolver struct
type PostResolver struct {
	post *store.Post
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

func (r *Resolver) CreateUser(ctx context.Context, args struct{ Input CreateUserInput }) (*UserResolver, error) {

    hashedPassword, er := auth.HashPassword(args.Input.Password)
	if er != nil {
        log.Fatalln(er);
	}
    
    newUser := store.User{
		FirstName: args.Input.FirstName,
		LastName:  args.Input.LastName,
		Email:     args.Input.Email,
		Username:  args.Input.Username,
		Password:  hashedPassword,
		Gender:    args.Input.Gender,
		Token:     args.Input.Token,
	}

	var id int32
	err := db.QueryRow(
		"INSERT INTO users (first_name, last_name, username, email, password, gender, token) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		newUser.FirstName, newUser.LastName, newUser.Username, newUser.Email, newUser.Password, newUser.Gender, newUser.Token,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	newUser.ID = id
    secret := []byte(os.Getenv("JWT_SECRET_TOKEN"));
	token, err := auth.CreateJWT(secret, int(newUser.ID));
    if err != nil {
        log.Fatal(err);
    }

	newUser.Token = token
    _, err = db.Exec("UPDATE users SET token = $1 WHERE id = $2", token, id)
    if err != nil {
		return nil, err
	}
    utils.SendEmail("spamsama91@gmail.com");
	

	return &UserResolver{user: &newUser}, nil
}

func (r *Resolver) Posts(ctx context.Context) ([]*PostResolver, error) {
	var posts []store.Post
	err := db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	postResolvers := make([]*PostResolver, len(posts))
	for i := range posts {
		postResolvers[i] = &PostResolver{post: &posts[i]}
	}
	return postResolvers, nil
}

func (r *Resolver) Post(ctx context.Context, args struct{ ID int32 }) (*PostResolver, error) {
	var post store.Post
	err := db.Get(&post, "SELECT * FROM posts WHERE id=$1", args.ID)
	if err != nil {
		return nil, err
	}
	return &PostResolver{post: &post}, nil
}

// Field resolvers for User
func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.user.ID)))
}

func (r *UserResolver) FirstName() string {
	return r.user.FirstName
}

func (r *UserResolver) LastName() string {
	return r.user.LastName
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

func (r *UserResolver) Gender() string {
	return r.user.Gender
}

func (r *UserResolver) Token() string {
	return r.user.Token
}

// Field resolvers for Post
func (r *PostResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.post.ID)))
}

func (r *PostResolver) Title() string {
	return r.post.Title
}

func (r *PostResolver) Content() string {
	return r.post.Content
}

func (r *PostResolver) UserID() graphql.ID {
	return graphql.ID(strconv.Itoa(r.post.UserID))
}
