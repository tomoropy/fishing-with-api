package infra

import (
	"context"
	// "database/sql"

	"github.com/tomoropy/fishing-with-api/domain/model"
	"github.com/tomoropy/fishing-with-api/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) SelectAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User

	result := ur.DB.Find(&users)
	err := result.Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) SelectUserByID(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	result := ur.DB.First(&user, "id = ?", id)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) CreateUser(
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

	result := ur.DB.Create(&user)
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

	result := ur.DB.First(&user, "id = ?", id)
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

	result = ur.DB.Save(&user)
	err = result.Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id int) error {
	var user model.User

	result := ur.DB.Where("id = ?", id).Delete(&user)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

// // invitaions
// type invRepository struct {
// 	DB *gorm.DB
// }

// func NewInvRepostitory(db *gorm.DB) repository.IInvRepository {
// 	return &invRepository{
// 		DB: db,
// 	}
// }

// // photo
// type photoRepository struct {
// 	DB *gorm.DB
// }

// func NewPhotoRepory(db *gorm.DB) repository.IPhotoRepository {
// 	return &photoRepository{
// 		DB: db,
// 	}
// }
