package repository

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
)

// I（interface） UserRepository
type IUserRepository interface {
	SelectAllUser(ctx context.Context) ([]model.User, error)
	SelectUserByID(ctx context.Context, id int) (*model.User, error)
}