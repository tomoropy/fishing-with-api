package repository

import (
	"context"

	"github.com/tomoropy/clean-arc-go/domain/model"
)

// I（interface） UserRepository
type IUserRepository interface {
	SelectAllUser(ctx context.Context) ([]model.User, error)
	SelectUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, username string, email string, password string, profileText string, profileImage string, profileHeader string) (*model.User, error)
	UpdateUser(ctx context.Context, id int, username string, email string, password string, profileText string, profileImage string, profileHeader string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) (bool, error)
}

type IInvRepository interface {
	SelectAllInvitation(ctx context.Context) ([]model.Invitation, error)
	SelectInvitationByUserID(ctx context.Context, userID int) ([]model.Invitation, error)
	CreateInvitation(ctx context.Context, userID int, comment string, place string, startTime string, endTime string) (*model.Invitation, error)
	UpdateInvitation(ctx context.Context, id, int, userID int, comment string, place string, startTime string, endTime string) (*model.Invitation, error)
	DeleteInvitation(ctx context.Context, id, int, userID int) error
}

type IPhotoRepository interface {
	SelectAllPhotoByInvitationID(ctx context.Context, invID int) (*model.Photo, error)
	CreatePhotoByInvitationID(ctx context.Context, invID int, userID int, image string) (*model.Photo, error)
	DeletePhotoByInvitationID(ctx context.Context, invID int, userID int) error
}
