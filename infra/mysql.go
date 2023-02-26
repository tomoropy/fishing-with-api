package infra

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/domain/repository"
)

// ---------------------------------------------------------------------------------------------------------------------------------------
// user repository
type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------

func (ur *userRepository) SelectAllUser(ctx context.Context) ([]entity.User, error) {

	var users []entity.User

	findAllUserSql := "SELECT * FROM users"

	err := ur.db.Select(&users, findAllUserSql)
	if err != nil {
		log.Error("failed to select all user" + err.Error())
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) SelectUserByUID(ctx context.Context, uid string) (*entity.User, error) {

	var user entity.User

	findUserByUIDSql := "SELECT * FROM users WHERE uid = ?"

	err := ur.db.Get(&user, findUserByUIDSql, uid)
	if err != nil {
		log.Error("failed to select user by uid" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) SelectUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	var user entity.User

	findUserByEmailSql := "SELECT * FROM users WHERE email = ?"

	err := ur.db.Get(&user, findUserByEmailSql, email)
	if err != nil {
		log.Error("failed to select user by email" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) InsertUser(ctx context.Context, user entity.User) (*entity.User, error) {

	insertUserSql := "INSERT INTO users (uid, username, email, hashed_password, text, avater, header, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := ur.db.Exec(insertUserSql, user.UID, user.Username, user.Email, user.HashedPassword, user.Text, user.Avater, user.CreatedAt, user.Header)
	if err != nil {
		log.Error("failed to insert user" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user entity.User) (*entity.User, error) {

	updateUserSql := "UPDATE users SET username = ?, email = ?, hashed_password = ?, text = ?, avater = ?, header = ? WHERE uid = ?"

	_, err := ur.db.Exec(updateUserSql, user.Username, user.Email, user.HashedPassword, user.Text, user.Avater, user.Header, user.UID)
	if err != nil {
		log.Error("failed to update user" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, uid string) error {

	deleteUserSql := "DELETE FROM users WHERE uid = ?"

	_, err := ur.db.Exec(deleteUserSql, uid)
	if err != nil {
		log.Error("failed to delete user" + err.Error())
		return err
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------

// invitaion repository
// type invRepository struct {
// 	db *gorm.DB
// }

// func NewInvRepostitory(db *gorm.DB) repository.InvRepository {
// 	return &invRepository{
// 		db: db,
// 	}
// }

// func (ir invRepository) SelectInv(ctx context.Context, id int) (*entity.Invitation, error) {
// 	var inv entity.Invitation

// 	result := ir.db.First(&inv, "id = ?", id)
// 	err := result.Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &inv, nil
// }

// func (ir invRepository) SelectAllInvitation(ctx context.Context) ([]entity.Invitation, error) {
// 	var inv []entity.Invitation

// 	result := ir.db.Find(&inv)
// 	err := result.Error

// 	if err != nil {
// 		return nil, err
// 	}
// 	return inv, nil
// }

// func (ir invRepository) InsertInvitation(ctx context.Context, userID int, comment string, place string) (*entity.Invitation, error) {
// 	inv := entity.Invitation{
// 		UserID:  userID,
// 		Comment: comment,
// 		Place:   place,
// 	}

// 	result := ir.db.Create(&inv)
// 	err := result.Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &inv, nil
// }

// func (ir invRepository) SelectInvitationByUserID(ctx context.Context, userID int) ([]entity.Invitation, error) {
// 	var invs []entity.Invitation

// 	result := ir.db.Where("user_id = ?", userID).Find(&invs)
// 	err := result.Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return invs, nil
// }

// func (ir invRepository) UpdateInvitation(ctx context.Context, id int, comment string, place string) (*entity.Invitation, error) {
// 	var inv entity.Invitation

// 	result := ir.db.First(&inv, "id = ?", id)
// 	err := result.Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	inv.Comment = comment
// 	inv.Place = place

// 	result = ir.db.Save(&inv)
// 	err = result.Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &inv, nil
// }

// func (ir invRepository) DeleteInvitation(ctx context.Context, id int) error {
// 	var inv entity.Invitation

// 	result := ir.db.Where("id = ?", id).Delete(&inv)
// 	err := result.Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
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
