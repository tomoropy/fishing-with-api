package infra

import (
	"context"

	"github.com/tomoropy/fishing-with-api/domain/model"
	"github.com/tomoropy/fishing-with-api/domain/repository"
	"gorm.io/gorm"
)

// user repository
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) SelectAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User

	result := ur.db.Find(&users)
	err := result.Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) SelectUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	result := ur.db.First(&user, "username = ?", username)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) SelectUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	result := ur.db.First(&user, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) InsertUser(
	ctx context.Context,
	username string,
	email string,
	password string,
	text string,
	avater string,
	header string,
) (*model.User, error) {
	user := model.User{
		Username:       username,
		Email:          email,
		HashedPassword: password,
		Text:           text,
		Avater:         avater,
		Header:         header,
	}

	result := ur.db.Create(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) UpdateUser(
	ctx context.Context,
	id int,
	username string,
	email string,
	password string,
	text string,
	avater string,
	header string,
) (*model.User, error) {
	var user model.User

	result := ur.db.First(&user, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	}

	user.Username = username
	user.Email = email
	user.HashedPassword = password
	user.Text = text
	user.Avater = avater
	user.Header = header

	result = ur.db.Save(&user)
	err = result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id int) error {
	var user model.User

	result := ur.db.Where("id = ?", id).Delete(&user)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

// invitaion repository
type invRepository struct {
	db *gorm.DB
}

func NewInvRepostitory(db *gorm.DB) repository.InvRepository {
	return &invRepository{
		db: db,
	}
}

func (ir invRepository) SelectInv(ctx context.Context, id int) (*model.Invitation, error) {
	var inv model.Invitation

	result := ir.db.First(&inv, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &inv, nil
}

func (ir invRepository) SelectAllInvitation(ctx context.Context) ([]model.Invitation, error) {
	var inv []model.Invitation

	result := ir.db.Find(&inv)
	err := result.Error

	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (ir invRepository) InsertInvitation(ctx context.Context, userID int, comment string, place string) (*model.Invitation, error) {
	inv := model.Invitation{
		UserID:  userID,
		Comment: comment,
		Place:   place,
	}

	result := ir.db.Create(&inv)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

func (ir invRepository) SelectInvitationByUserID(ctx context.Context, userID int) ([]model.Invitation, error) {
	var invs []model.Invitation

	result := ir.db.Where("user_id = ?", userID).Find(&invs)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return invs, nil
}

func (ir invRepository) UpdateInvitation(ctx context.Context, id int, comment string, place string) (*model.Invitation, error) {
	var inv model.Invitation

	result := ir.db.First(&inv, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	}

	inv.Comment = comment
	inv.Place = place

	result = ir.db.Save(&inv)
	err = result.Error
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

func (ir invRepository) DeleteInvitation(ctx context.Context, id int) error {
	var inv model.Invitation

	result := ir.db.Where("id = ?", id).Delete(&inv)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

// // photo
// type photoRepository struct {
// 	DB *gorm.DB
// }

// func NewPhotoRepory(db *gorm.DB) repository.IPhotoRepository {
// 	return &photoRepository{
// 		DB: db,
// 	}
// }
