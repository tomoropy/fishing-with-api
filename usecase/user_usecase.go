package usecase

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
	"github.com/tomoropy/clean-arc-go/domain/repository"
)

type IUserUsecase interface {
	FindAllUser(ctx context.Context) ([]model.User, error)
	FindUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, username string, email string, password string, age int) (*model.User, error)
	UpdateUser(ctx context.Context, id int, username string, email string, password string, age int) (*model.User, error)
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

func (uu *userUsecase) CreateUser(ctx context.Context, username string, email string, password string, age int) (*model.User, error) {
	createdUser, err := uu.ur.CreateUser(ctx, username, email, password, age)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (uu *userUsecase) UpdateUser(ctx context.Context, id int, username string, email string, password string, age int) (*model.User, error) {
	updatedUser, err := uu.ur.UpdateUser(ctx, id, username, email, password, age)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
