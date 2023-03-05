package repository

import (
	"context"

	"github.com/tomoropy/fishing-with-api/domain/entity"
)

type UserRepository interface {
	SelectAllUser(ctx context.Context) ([]entity.User, error)
	SelectUserByUID(ctx context.Context, uid string) (*entity.User, error)
	SelectUserByEmail(ctx context.Context, email string) (*entity.User, error)
	InsertUser(ctx context.Context, user entity.User) (*entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, uid string) error
}

// type TweetRepository interface {
// 	SelectAllTweet(ctx context.Context) ([]entity.Tweet, error)
// 	SelectTweetByUID(ctx context.Context, uid string) (*entity.Tweet, error)
// 	SelectTweetByUserID(ctx context.Context, userID string) (*entity.Tweet, error)
// 	InsertTweet(ctx context.Context, tweet entity.Tweet) (*entity.Tweet, error)
// 	UpdateTweet(ctx context.Context, tweet entity.Tweet) (*entity.Tweet, error)
// 	DeleteTweet(ctx context.Context, uid string) error
// }
