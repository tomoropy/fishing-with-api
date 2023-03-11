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
	SaveUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, uid string) error

	// tweet
	SaveTweet(ctx context.Context, tweet *entity.Tweet) (*entity.Tweet, error)
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
	user, err := qs.ur.SelectByEmail(ctx, email)
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
	users, err := qs.ur.SelectAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (qs *queyrService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := qs.ur.SelectByUID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (qs *queyrService) ListTweets(ctx context.Context) ([]entity.Tweet, error) {
	tweets, err := qs.tr.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (qs *queyrService) GetTweet(ctx context.Context, id string) (*entity.Tweet, error) {
	tweet, err := qs.tr.SelectByUID(ctx, id)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (qs *queyrService) GetTweetsByUserUID(ctx context.Context, userID string) ([]entity.Tweet, error) {
	tweets, err := qs.tr.SelectByUserUID(ctx, userID)
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

func (ms *mutationService) SaveUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	// password hash
	hashedPassword, err := util.HashPassword(user.HashedPassword)
	if err != nil {
		return nil, err
	}
	user.HashedPassword = hashedPassword

	selectUser, _ := ms.ur.SelectByUID(ctx, user.UID)

	if selectUser == nil {
		// create user
		if _, err := ms.ur.Insert(ctx, *user); err != nil {
			return nil, err
		}

	} else {
		// update user
		if _, err := ms.ur.Update(ctx, *user); err != nil {
			return nil, err
		}
	}

	user, err = ms.ur.SelectByUID(ctx, user.UID)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (ms *mutationService) DeleteUser(ctx context.Context, uid string) error {
	err := ms.ur.DeleteTX(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}

func (ms *mutationService) SaveTweet(ctx context.Context, tweet *entity.Tweet) (*entity.Tweet, error) {
	selectTweet, _ := ms.tr.SelectByUID(ctx, tweet.UID)

	if selectTweet == nil {
		// create tweet
		if _, err := ms.tr.Insert(ctx, *tweet); err != nil {
			return nil, err
		}

	} else {
		// update tweet
		if _, err := ms.tr.Update(ctx, *tweet); err != nil {
			return nil, err
		}
	}

	tweet, err := ms.tr.SelectByUID(ctx, tweet.UID)
	if err != nil {
		return nil, err
	}

	return tweet, nil

}

func (ms *mutationService) DeleteTweet(ctx context.Context, uid string) error {
	err := ms.tr.Delete(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}
