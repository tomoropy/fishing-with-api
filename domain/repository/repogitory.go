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

// type InvRepository interface {
// 	SelectAllInv(ctx context.Context) ([]entity.Invitation, error)
// 	SelectInvByUID(ctx context.Context, uid string) (*entity.Invitation, error)
// 	InsertInv(ctx context.Context, inv entity.Invitation) (*entity.Invitation, error)
// 	UpdateInv(ctx context.Context, inv entity.Invitation) (*entity.Invitation, error)
// 	DeleteInv(ctx context.Context, uid string) error
// }
