package resolvers

import "github.com/nagisa599/go-graphql-template/internal/handler"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserHandler handler.UserHandler
}

func NewResolver(userHandler handler.UserHandler) *Resolver {
	return &Resolver{
		UserHandler: userHandler,
	}
}