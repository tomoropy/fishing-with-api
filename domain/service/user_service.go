package service

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
	"github.com/tomoropy/clean-arc-go/domain/repository"
)

type IUserService interface {
	FindAllUser(ctx context.Context) ([]model.User, error)
	FindUserByID(ctx context.Context, id int) (*model.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(su repository.IUserRepository) IUserService {
	return &userService{
		repo: su,
	}
}

func (su *userService) FindAllUser(ctx context.Context) ([]model.User, error) {
	return su.repo.SelectAllUser(ctx)
}

func (su *userService) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	return su.repo.SelectUserByID(ctx, id)
}
