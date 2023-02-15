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

func (ur *userRepository) SelectAllUser(ctx context.Context) ([]model.User, error) {
	var users []model.User

	rows, err := ur.DB.Query("SELECT * FROM user")
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

func (ur *userRepository) SelectUserByID(ctx context.Context, sutudentID int) (*model.User, error) {
	var user model.User

	row := ur.DB.QueryRow("SELECT * FROM user WHERE id = ?", sutudentID)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Age); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	return &user, nil
}

func (ur *userRepository) CreateUser(ctx context.Context, username string, email string, password string, age int) (*model.User, error) {

	result, err := ur.DB.Exec("INSERT INTO user (username, email, password, age) VALUES (?, ?, ?, ?);", username, email, password, age)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var user model.User
	user.ID = int(id)
	user.Email = email
	user.Username = username
	user.Password = password
	user.Age = age

	return &user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, id int, username string, email string, password string, age int) (*model.User, error) {

	_, err := ur.DB.Exec("UPDATE user SET username = ?, email = ?, password = ?, age = ? WHERE id = ?", username, email, password, age, id)
	if err != nil {
		return nil, err
	}

	var user model.User
	user.ID = id
	user.Email = email
	user.Username = username
	user.Password = password
	user.Age = age

	return &user, nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id int) (bool, error) {
	_, err := ur.DB.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return false, err
	}

	return true, nil
}
