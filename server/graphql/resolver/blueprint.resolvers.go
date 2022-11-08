package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"blueprint/ent"
	"blueprint/graphql/generated"
	"blueprint/graphql/model"
	"context"

	"github.com/rs/xid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	user, err := r.client.User.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, userID xid.ID, input ent.CreatePostInput) (*ent.Post, error) {
	client := ent.FromContext(ctx)
	user, err := client.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	post, err := client.Post.Create().
		SetContent(input.Content).
		SetAuthor(user).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// BulkUpsertUser is the resolver for the bulkUpsertUser field.
func (r *mutationResolver) BulkUpsertUser(ctx context.Context, input []*model.BulkUpsertUserInput) (bool, error) {
	userCreates := make([]*ent.UserCreate, len(input))
	for i := range input {
		c := r.client.User.Create()
		if input[i].ID != nil {
			c.SetID(*input[i].ID)
		}
		c.SetName(input[i].Name)
		c.SetAge(input[i].Age)
		userCreates[i] = c
	}
	if err := r.client.User.CreateBulk(userCreates...).OnConflict().UpdateNewValues().Exec(ctx); err != nil {
		return false, err
	}

	return true, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id xid.ID) (*ent.User, error) {
	user, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id xid.ID) (*ent.Post, error) {
	post, err := r.client.Post.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
