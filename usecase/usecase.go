package usecase

import (
	// "context"
	// "errors"
	// "log"
	// "time"

	// "github.com/tomoropy/fishing-with-api/auth/token"
	// "github.com/tomoropy/fishing-with-api/config"
	// "github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/domain/repository"
	// "github.com/tomoropy/fishing-with-api/util"
)

type Usecase interface {
	// user
	// Login(ctx context.Context, email string, password string) (*entity.User, string, error)
	// Register(ctx context.Context, username string, email string, password string, text string, avater string, header string) (*entity.User, error)
	// FindAllUser(ctx context.Context) ([]entity.User, error)
	// FindUserByID(ctx context.Context, id int) (*entity.User, error)
	// UpdateUser(ctx context.Context, id int, username string, email string, password string, text string, avater string, header string) (*entity.User, error)
	// DeleteUser(ctx context.Context, id int) error

	// // invtation
	// FindInv(ctx context.Context, id int) (*entity.Invitation, error)
	// FindAllInv(ctx context.Context) ([]entity.Invitation, error)
	// CreateInv(ctx context.Context, userID int, comment string, place string) (*entity.Invitation, error)
	// FindInvitationByUserID(ctx context.Context, userID int) ([]entity.Invitation, error)
	// UpdateInv(ctx context.Context, id int, comment string, place string) (*entity.Invitation, error)
	// DeleteInv(ctx context.Context, id int) error
}

type usecase struct {
	ur repository.UserRepository
	// ir repository.InvRepository
}

func NewUsecase(
	ur repository.UserRepository,
	// ir repository.InvRepository
) Usecase {
	return &usecase{
		ur: ur,
		// ir: ir,
	}
}

// func (u *usecase) Login(ctx context.Context, username string, password string) (*entity.User, string, error) {
// user, err := u.ur.SelectUserByUsername(ctx, username)
// if err != nil {
// 	return nil, "", err
// }

// 	if err = util.CheckPassword(password, user.HashedPassword); err != nil {
// 		return nil, "", errors.New("password is not correct")
// 	}

// 	// load config
// 	config, err := config.Load()
// 	if err != nil {
// 		log.Fatalf("failed to load config: %v", err)
// 	}

// 	tokenMaker, err := token.NewJWTMaker(config.Auth.SecretKey)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	token, _, err := tokenMaker.CreateTocken(username, 24*time.Hour)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return user, token, nil
// }

// func (u *usecase) Register(ctx context.Context, username string, email string, password string, text string, avater string, header string) (*entity.User, error) {
// 	hashedPassword, err := util.HashPassword(password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	createdUser, err := u.ur.InsertUser(ctx, username, email, hashedPassword, text, avater, header)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return createdUser, nil
// }

// func (u *usecase) FindAllUser(ctx context.Context) ([]entity.User, error) {
// 	users, err := u.ur.SelectAllUser(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func (u *usecase) FindUserByID(ctx context.Context, id int) (*entity.User, error) {
// 	user, err := u.ur.SelectUserByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func (u *usecase) UpdateUser(ctx context.Context, id int, username string, email string, password string, text string, avater string, header string) (*entity.User, error) {
// 	updatedUser, err := u.ur.UpdateUser(ctx, id, username, email, password, text, avater, header)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return updatedUser, nil
// }

// func (u *usecase) DeleteUser(ctx context.Context, id int) error {
// 	err := u.ur.DeleteUser(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // invitations usecase
// func (u *usecase) FindInv(ctx context.Context, id int) (*entity.Invitation, error) {
// 	inv, err := u.ir.SelectInv(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return inv, nil
// }

// func (u *usecase) FindAllInv(ctx context.Context) ([]entity.Invitation, error) {
// 	invs, err := u.ir.SelectAllInvitation(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return invs, nil
// }

// func (u *usecase) CreateInv(ctx context.Context, userID int, comment string, place string) (*entity.Invitation, error) {
// 	inv, err := u.ir.InsertInvitation(ctx, userID, comment, place)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return inv, nil
// }

// func (u *usecase) FindInvitationByUserID(ctx context.Context, userID int) ([]entity.Invitation, error) {
// 	invs, err := u.ir.SelectInvitationByUserID(ctx, userID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return invs, nil
// }

// func (u *usecase) UpdateInv(ctx context.Context, id int, comment string, place string) (*entity.Invitation, error) {
// 	inv, err := u.ir.UpdateInvitation(ctx, id, comment, place)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return inv, nil
// }

// func (u *usecase) DeleteInv(ctx context.Context, id int) error {
// 	err := u.ir.DeleteInvitation(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// type IinvUsecase interface {
// 	FindInv(ctx context.Context, id int) *entity.Invitation
// 	AllInv(ctx context.Context) []entity.Invitation
// 	UserInv(ctx context.Context, userID int) []entity.Invitation
// 	CreateInv(ctx context.Context, userID int) *entity.Invitation
// 	UpdateInv(ctx context.Context, userID int) *entity.Invitation
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
