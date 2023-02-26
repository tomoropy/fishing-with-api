package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/tomoropy/fishing-with-api/auth/token"
	"github.com/tomoropy/fishing-with-api/config"
	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/domain/repository"
	"github.com/tomoropy/fishing-with-api/util"
)

// ---------------------------------------------------------------------------------------------------------------------------------------
// interface

type QueryService interface {
	Login(ctx context.Context, email string, password string) (*entity.User, string, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
}

type MutationService interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// query service

type queyrService struct {
	ur repository.UserRepository
}

// constructor
func NewQueryService(
	ur repository.UserRepository,
) QueryService {

	return &queyrService{
		ur: ur,
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// QueryServiceの実装

func (qs *queyrService) Login(ctx context.Context, email string, password string) (*entity.User, string, error) {
	user, err := qs.ur.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, "", err
	}

	// passwrod check
	if err = util.CheckPassword(password, user.HashedPassword); err != nil {
		return nil, "", errors.New("password is not correct")
	}

	// load config
	config, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	tokenMaker, err := token.NewJWTMaker(config.Auth.SecretKey)
	if err != nil {
		return nil, "", err
	}

	// create token
	token, _, err := tokenMaker.CreateTocken(email, 24*time.Hour)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (qs *queyrService) ListUsers(ctx context.Context) ([]entity.User, error) {
	users, err := qs.ur.SelectAllUser(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (qs *queyrService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := qs.ur.SelectUserByUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// mutation service

type mutationService struct {
	ur repository.UserRepository
}

// constructor
func NewMutationService(
	ur repository.UserRepository,
) MutationService {

	return &mutationService{
		ur: ur,
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// MutationServiceの実装

func (ms *mutationService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	// password hash
	hashedPassword, err := util.HashPassword(user.HashedPassword)
	if err != nil {
		return nil, err
	}
	user.HashedPassword = hashedPassword

	// create user
	createdUser, err := ms.ur.InsertUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (ms *mutationService) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	// password hash
	hashedPassword, err := util.HashPassword(user.HashedPassword)
	if err != nil {
		return nil, err
	}
	user.HashedPassword = hashedPassword

	// update user
	updatedUser, err := ms.ur.UpdateUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (ms *mutationService) DeleteUser(ctx context.Context, id string) error {
	err := ms.ur.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
