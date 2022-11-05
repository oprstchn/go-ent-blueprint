package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"blueprint/ent"
	"blueprint/graphql/generated"
	"context"

	"github.com/rs/xid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id xid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []xid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	return r.client.Post.Query().Paginate(ctx, after, first, before, last, ent.WithPostFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	users, err := r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
