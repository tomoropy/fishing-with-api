package usecase

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
	"github.com/tomoropy/clean-arc-go/domain/service"
)

type IUserUsecase interface {
	FindAllUser(ctx context.Context) ([]model.User, error)
	FindUserByID(ctx context.Context, id int) (*model.User, error)
}
type userUsecase struct {
	svc service.IUserService
}

func NewUserUsecase(su service.IUserService) IUserUsecase {
	return &userUsecase{
		svc: su,
	}
}

func (uu *userUsecase) FindAllUser(ctx context.Context) ([]model.User, error) {
	users, err := uu.svc.FindAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *userUsecase) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := uu.svc.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
