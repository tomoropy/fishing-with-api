package infra

import (
	"context"
	"database/sql"

	"github.com/tomoropy/clean-arc-go/domain/model"
	"github.com/tomoropy/clean-arc-go/domain/repository"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (sr *userRepository) SelectAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User

	rows, err := sr.DB.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return users, nil
}

func (sr *userRepository) SelectUserByID(ctx context.Context, sutudentID int) (*model.User, error) {
	var user model.User

	row := sr.DB.QueryRow("SELECT * FROM user WHERE id = ?", sutudentID)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}

	return &user, nil
}
