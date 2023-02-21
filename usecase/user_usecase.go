package usecase

import (
	"context"

	"github.com/tomoropy/fishing-with-api/domain/model"
	"github.com/tomoropy/fishing-with-api/domain/repository"
)

type Usecase interface {
	// user
	FindAllUser(ctx context.Context) ([]model.User, error)
	FindUserByID(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, username string, email string, password string, text string, avater string, header string) (*model.User, error)
	UpdateUser(ctx context.Context, id int, username string, email string, password string, text string, avater string, header string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error

	// invtation
	FindAllInv(ctx context.Context) ([]model.Invitation, error)
}

type usecase struct {
	ur repository.UserRepository
	ir repository.InvRepository
}

func NewUsecase(ur repository.UserRepository, ir repository.InvRepository) Usecase {
	return &usecase{
		ur: ur,
		ir: ir,
	}
}

func (u *usecase) FindAllUser(ctx context.Context) ([]model.User, error) {
	users, err := u.ur.SelectAllUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *usecase) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	user, err := u.ur.SelectUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usecase) CreateUser(ctx context.Context, username string, email string, password string, text string, avater string, header string) (*model.User, error) {
	createdUser, err := u.ur.InsertUser(ctx, username, email, password, text, avater, header)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (u *usecase) UpdateUser(ctx context.Context, id int, username string, email string, password string, text string, avater string, header string) (*model.User, error) {
	updatedUser, err := u.ur.UpdateUser(ctx, id, username, email, password, text, avater, header)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *usecase) DeleteUser(ctx context.Context, id int) error {
	err := u.ur.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// invitations usecase
func (u *usecase) FindAllInv(ctx context.Context) ([]model.Invitation, error) {
	invs, err := u.ir.SelectAllInvitation(ctx)
	if err != nil {
		return nil, err
	}
	return invs, nil
}

// type IinvUsecase interface {
// 	FindInv(ctx context.Context, id int) *model.Invitation
// 	AllInv(ctx context.Context) []model.Invitation
// 	UserInv(ctx context.Context, userID int) []model.Invitation
// 	CreateInv(ctx context.Context, userID int) *model.Invitation
// 	UpdateInv(ctx context.Context, userID int) *model.Invitation
// 	DeleteInv(ctx context.Context, userID int) error
// }

// type invUsecase struct {
// 	repo repository.IInvRepository
// }

// func newInvUsecase(ir repository.IInvRepository) invUsecase {
// 	return &invUsecase{
// 		ir: ir,
// 	}
// }
