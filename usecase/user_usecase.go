package usecase

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
	"github.com/tomoropy/clean-arc-go/domain/repository"
)

type IUserUsecase interface {
	FindAllUser(ctx context.Context) ([]model.User, error)
	FindUserByID(ctx context.Context, id int) (*model.User, error)
}
type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		ur: ur,
	}
}

func (uu *userUsecase) FindAllUser(ctx context.Context) ([]model.User, error) {
	users, err := uu.ur.SelectAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *userUsecase) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := uu.ur.SelectUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
