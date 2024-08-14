package repository

import (
	"context"

	"github.com/nagisa599/go-graphql-template/internal/domain"
)
type UserEntity struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Email    string
	Password string
}

type UserRepository struct {
	dh DatabaseHandler
}

func NewUserRepository(databaseHandler DatabaseHandler) *UserRepository {
	return &UserRepository{dh: databaseHandler}
}

func newCarEntity(user domain.User) *UserEntity {
	return &UserEntity {
		ID: uint(user.ID),
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
}
func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	newUser := newCarEntity(*user)
	if err := ur.dh.Conn(ctx).Table("users").Create(newUser).Error; err != nil {
		return err
	}
	return nil
}
func (ur *UserRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	var users []*UserEntity
	if err := ur.dh.Conn(ctx).Table("users").Find(&users).Error; err != nil {
		return nil, err
	}
	var domainUsers []*domain.User
	for _, user := range users {
		domainUsers = append(domainUsers, &domain.User{
			ID:       int(user.ID),
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return domainUsers, nil
}