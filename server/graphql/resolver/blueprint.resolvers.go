package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"blueprint/ent"
	"blueprint/graphql/generated"
	"context"
	"fmt"

	"github.com/rs/xid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id xid.ID) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
