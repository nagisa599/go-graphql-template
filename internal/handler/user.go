package handler

import (
	"context"

	graphql_model "github.com/nagisa599/go-graphql-template/graphql/model"
	"github.com/nagisa599/go-graphql-template/internal/domain"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, input *domain.User) error
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}
type UserHandler struct {
  uu UserUsecase
}

func NewUserHandler(userUsecase UserUsecase) *UserHandler {
	return &UserHandler{uu: userUsecase}
}

func (uh *UserHandler) GetAllUsers(ctx context.Context) ([]*graphql_model.User, error) {
	users,err := uh.uu.GetAllUsers(ctx)
	if (err != nil) {
		return nil, err
	}
	var gUsers []*graphql_model.User
	for _, users := range users {
		gUsers = append(gUsers, &graphql_model.User{
			ID: string(users.ID),
			Name: users.Name,
			Email: users.Email,
			Password: users.Password,
		})
	}
	return gUsers, nil
    
}

func (uh *UserHandler) CreateUser(ctx context.Context, input graphql_model.NewUser) (*graphql_model.ResponseStatus, error) {
	// ドメインの実体を作成する
	user := &domain.User {
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}
	err := uh.uu.CreateUser(ctx, user)
	if (err != nil) {
		return nil, err
	}
	return &graphql_model.ResponseStatus{
		Status: "success",
	}, nil
}