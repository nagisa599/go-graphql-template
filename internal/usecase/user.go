package usecase

import (
	"context"

	"github.com/nagisa599/go-graphql-template/internal/domain"
)

type UserUsecase struct {
}

func NewUserUsecase() *UserUsecase {
	return &UserUsecase{}
}

func (u *UserUsecase) CreateUser(ctx context.Context, input *domain.User) error {
	// ドメインの実体を作成する
	// ここに実装を書く
	return nil
}
