package repository

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
)

// I（interface） UserRepository
type IUserRepository interface {
	SelectAllUser(ctx context.Context) ([]model.User, error)
	SelectUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, username string, email string, password string, age int) (*model.User, error)
	UpdateUser(ctx context.Context, id int, username string, email string, password string, age int) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}
