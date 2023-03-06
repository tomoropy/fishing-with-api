package service

import (
	"context"
	"errors"

	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/domain/repository"
	"github.com/tomoropy/fishing-with-api/util"
)

// ---------------------------------------------------------------------------------------------------------------------------------------
// interface

type QueryService interface {
	// user
	Login(ctx context.Context, email string, password string) (*entity.User, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)

	// tweet
	ListTweets(ctx context.Context) ([]entity.Tweet, error)
	GetTweet(ctx context.Context, id string) (*entity.Tweet, error)
	GetTweetsByUserUID(ctx context.Context, userID string) ([]entity.Tweet, error)
}

type MutationService interface {
	// user
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, uid string) error

	// tweet
	CreateTweet(ctx context.Context, tweet *entity.Tweet) (*entity.Tweet, error)
	UpdateTweet(ctx context.Context, tweet *entity.Tweet) (*entity.Tweet, error)
	DeleteTweet(ctx context.Context, uid string) error
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// query service

type queyrService struct {
	ur repository.UserRepository
	tr repository.TweetRepository
}

// constructor
func NewQueryService(
	ur repository.UserRepository,
	tr repository.TweetRepository,
) QueryService {

	return &queyrService{
		ur: ur,
		tr: tr,
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// QueryServiceの実装

func (qs *queyrService) Login(ctx context.Context, email string, password string) (*entity.User, error) {
	user, err := qs.ur.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// passwrod check
	if err = util.CheckPassword(password, user.HashedPassword); err != nil {
		return nil, errors.New("password is not correct")
	}

	return user, nil
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

func (qs *queyrService) ListTweets(ctx context.Context) ([]entity.Tweet, error) {
	tweets, err := qs.tr.SelectAllTweet(ctx)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (qs *queyrService) GetTweet(ctx context.Context, id string) (*entity.Tweet, error) {
	tweet, err := qs.tr.SelectTweetByUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (qs *queyrService) GetTweetsByUserUID(ctx context.Context, userID string) ([]entity.Tweet, error) {
	tweets, err := qs.tr.SelectTweetByUserUID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// mutation service

type mutationService struct {
	ur repository.UserRepository
	tr repository.TweetRepository
}

// constructor
func NewMutationService(
	ur repository.UserRepository,
	tr repository.TweetRepository,
) MutationService {

	return &mutationService{
		ur: ur,
		tr: tr,
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

func (ms *mutationService) DeleteUser(ctx context.Context, uid string) error {
	err := ms.ur.DeleteUserTX(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}

func (ms *mutationService) CreateTweet(ctx context.Context, tweet *entity.Tweet) (*entity.Tweet, error) {
	createdTweet, err := ms.tr.InsertTweet(ctx, *tweet)
	if err != nil {
		return nil, err
	}

	return createdTweet, nil
}

func (ms *mutationService) UpdateTweet(ctx context.Context, tweet *entity.Tweet) (*entity.Tweet, error) {
	updatedTweet, err := ms.tr.UpdateTweet(ctx, *tweet)
	if err != nil {
		return nil, err
	}

	return updatedTweet, nil
}

func (ms *mutationService) DeleteTweet(ctx context.Context, uid string) error {
	err := ms.tr.DeleteTweet(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}
