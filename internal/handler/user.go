package handler

import (
	"context"

	graphql_model "github.com/nagisa599/go-graphql-template/graphql/model"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) CreateUser(ctx context.Context, input graphql_model.NewUser) (*graphql_model.ResponseStatus, error) {
	return &graphql_model.ResponseStatus{
		Status: "success",
	}, nil
}