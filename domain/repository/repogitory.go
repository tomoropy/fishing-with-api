package repository

import (
	"context"

	"github.com/tomoropy/fishing-with-api/domain/model"
)

// I（interface） UserRepository
type UserRepository interface {
	SelectAllUser(ctx context.Context) ([]model.User, error)
	SelectUserByID(ctx context.Context, id int) (*model.User, error)
	InsertUser(ctx context.Context, username string, email string, password string, text string, avater string, header string) (*model.User, error)
	UpdateUser(ctx context.Context, id int, username string, email string, password string, text string, avater string, header string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type InvRepository interface {
	SelectInv(ctx context.Context, id int) (*model.Invitation, error)
	SelectAllInvitation(ctx context.Context) ([]model.Invitation, error)
	SelectInvitationByUserID(ctx context.Context, userID int) ([]model.Invitation, error)
	InsertInvitation(ctx context.Context, userID int, comment string, place string) (*model.Invitation, error)
	UpdateInvitation(ctx context.Context, id int, comment string, place string) (*model.Invitation, error)
	DeleteInvitation(ctx context.Context, id int) error
}

// type IPhotoRepository interface {
// 	SelectAllPhotoByInvitationID(ctx context.Context, invID int) (*model.Photo, error)
// 	CreatePhotoByInvitationID(ctx context.Context, invID int, userID int, image string) (*model.Photo, error)
// 	DeletePhotoByInvitationID(ctx context.Context, invID int, userID int) error
// }
