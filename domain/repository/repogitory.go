package repository

import (
	"context"

	"github.com/tomoropy/fishing-with-api/domain/entity"
)

type UserRepository interface {
	SelectAll(ctx context.Context) ([]entity.User, error)
	SelectByUID(ctx context.Context, uid string) (*entity.User, error)
	SelectByEmail(ctx context.Context, email string) (*entity.User, error)
	Insert(ctx context.Context, user entity.User) (*entity.User, error)
	Update(ctx context.Context, user entity.User) (*entity.User, error)
	DeleteTX(ctx context.Context, uid string) error
}

type TweetRepository interface {
	SelectAll(ctx context.Context) ([]entity.Tweet, error)
	SelectByUID(ctx context.Context, uid string) (*entity.Tweet, error)
	SelectByUserUID(ctx context.Context, userID string) ([]entity.Tweet, error)
	Insert(ctx context.Context, tweet entity.Tweet) (*entity.Tweet, error)
	Update(ctx context.Context, tweet entity.Tweet) (*entity.Tweet, error)
	Delete(ctx context.Context, uid string) error
}
