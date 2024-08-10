package graphql_resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	graphql1 "github.com/nagisa599/go-graphql-template/graphql"
	graphql_model "github.com/nagisa599/go-graphql-template/graphql/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string, password string) (*graphql_model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) ([]*graphql_model.User, error) {
	panic(fmt.Errorf("not implemented: GetAllUsers - getAllUsers"))
}

// Todos is the resolver for the todos field.
func (r *userResolver) Todos(ctx context.Context, obj *graphql_model.User) ([]*graphql_model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
