package usecase

import (
	"context"

	"github.com/nagisa599/go-graphql-template/internal/domain"
)
type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
}

type UserUsecase struct {
	ur UserRepository
}

func NewUserUsecase(userRepository UserRepository) *UserUsecase {
	return &UserUsecase{ur : userRepository}
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	// ここに実装を書く
    users,err:= u.ur.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *UserUsecase) CreateUser(ctx context.Context, input *domain.User) error {
	// ドメインの実体を作成する
	if err := u.ur.CreateUser(ctx, input); err != nil {
		return err
	}
	// ここに実装を書く
	return nil
}
