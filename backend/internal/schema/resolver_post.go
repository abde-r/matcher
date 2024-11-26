package schema

import (
	"context"
	"matcher/internal/store"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

// PostResolver struct
type PostResolver struct {
	post *store.Post
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
