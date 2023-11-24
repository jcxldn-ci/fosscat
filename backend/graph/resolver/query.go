package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/structs"
)

// Checkout is the resolver for the checkout field.
func (r *queryResolver) Checkout(ctx context.Context) ([]*structs.Checkout, error) {
	// Get all checkouts. Proof of concept only, returns all fields!
	checkouts := []*structs.Checkout{}
	result := r.db.Find(&checkouts)
	return checkouts, result.Error
}

// Entity is the resolver for the entity field.
func (r *queryResolver) Entity(ctx context.Context) ([]*structs.Entity, error) {
	// Get all entities. Proof of concept only, returns all fields!
	entities := []*structs.Entity{}
	result := r.db.Find(&entities)
	return entities, result.Error
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context) ([]*structs.Item, error) {
	// Get all items. Proof of concept only, returns all fields!
	items := []*structs.Item{}
	result := r.db.Find(&items)
	return items, result.Error
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*structs.User, error) {
	// Get all users. Proof of concept only, returns all fields!
	users := []*structs.User{}
	result := r.db.Find(&users)
	return users, result.Error
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
