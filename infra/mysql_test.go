package infra_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/infra"
)

// UserRepository_Insert 正常系
func TestUserRepository_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.Background()
	dbx := sqlx.NewDb(db, "sqlmock")
	ur := infra.NewUserRepository(dbx)

	user := &entity.User{
		UID:            "uid",
		Username:       "username",
		Email:          "email",
		HashedPassword: "password",
		Text:           "text",
		Avater:         "avater",
		Header:         "header",
		CreatedAt:      "2020-01-01 00:00:00",
	}

	mock.ExpectExec(`INSERT INTO users`).
		WithArgs(user.UID, user.Username, user.Email, user.HashedPassword, user.Text, user.Avater, user.Header, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = ur.Insert(ctx, *user)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// UserRepository_Insert 異常系
func TestUserRepository_Insert_Invalid(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ctx := context.Background()
	dbx := sqlx.NewDb(db, "sqlmock")
	ur := infra.NewUserRepository(dbx)

	user := &entity.User{
		UID:            "",
		Username:       "",
		Email:          "",
		HashedPassword: "",
		Text:           "text",
		Avater:         "avater",
		Header:         "header",
		CreatedAt:      "invalid date format",
	}

	// Test case 1: Empty UID
	_, err = ur.Insert(ctx, *user)
	if err == nil {
		t.Errorf("expected error when UID is empty, but got nil")
	}

	// Test case 2: Empty username
	user.UID = "uid"
	_, err = ur.Insert(ctx, *user)
	if err == nil {
		t.Errorf("expected error when username is empty, but got nil")
	}

	// Test case 3: Empty email
	user.Username = "username"
	_, err = ur.Insert(ctx, *user)
	if err == nil {
		t.Errorf("expected error when email is empty, but got nil")
	}

	// Test case 4: Empty password
	user.Email = "email"
	_, err = ur.Insert(ctx, *user)
	if err == nil {
		t.Errorf("expected error when password is empty, but got nil")
	}

	// Test case 5: Invalid date format
	user.HashedPassword = "password"
	user.CreatedAt = "2020-01-01"
	_, err = ur.Insert(ctx, *user)
	if err == nil {
		t.Errorf("expected error when date format is invalid, but got nil")
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// func TestUserRepository_Insert(t *testing.T) {
// 	db, mock, err := sqlxmock.Newx()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// 正常系
// 	validUser := &entity.User{
// 		UID:            "uid",
// 		Username:       "username",
// 		Email:          "email",
// 		HashedPassword: "password",
// 		Text:           "text",
// 		Avater:         "avater",
// 		Header:         "header",
// 		CreatedAt:      "2020-01-01 00:00:00",
// 	}

// 	// 異常系(すでに登録されているユーザー)
// 	duplicateUser := &entity.User{
// 		UID:            "uid",
// 		Username:       "username",
// 		Email:          "email",
// 		HashedPassword: "password",
// 		Text:           "text",
// 		Avater:         "avater",
// 		Header:         "header",
// 		CreatedAt:      "2020-01-01 00:00:00",
// 	}

// 	// モック
// 	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (uid, username, email, hashed_password, text, avater, header, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")).
// 		WithArgs("uid", "username", "email", "password", "text", "avater", "header", "2020-01-01 00:00:00").
// 		WillReturnResult(sqlxmock.NewResult(1, 1))

// 	// リポジトリの生成
// 	repo := infra.NewUserRepository(db)

// 	// 正常系
// 	testUser, err := repo.Insert(context.Background(), *validUser)
// 	assert.NoError(t, err)
// 	assert.Equal(t, validUser, testUser)

// 	// 異常系
// 	_, err = repo.Insert(context.Background(), *duplicateUser)
// 	assert.Error(t, err)

// }

// func TestUserRepository_Update(t *testing.T) {
// 	db, mock, err := sqlxmock.Newx()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// 正常系
// 	validUser := &entity.User{
// 		UID:            "uid",
// 		Username:       "username",
// 		Email:          "email",
// 		HashedPassword: "password",
// 		Text:           "text",
// 		Avater:         "avater",
// 		Header:         "header",
// 		CreatedAt:      "2020-01-01 00:00:00",
// 	}

// 	// 異常系(存在しないユーザー)
// 	invalidUser := &entity.User{
// 		UID:            "invalid_uid",
// 		Username:       "username",
// 		Email:          "email",
// 		HashedPassword: "password",
// 		Text:           "text",
// 		Avater:         "avater",
// 		Header:         "header",
// 		CreatedAt:      "2020-01-01 00:00:00",
// 	}

// 	// モック
// 	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET username = ?, email = ?, hashed_password = ?, text = ?, avater = ?, header = ? WHERE uid = ?")).
// 		WithArgs("username", "email", "password", "text", "avater", "header", "uid").
// 		WillReturnResult(sqlxmock.NewResult(1, 1))

// 	// リポジトリの生成
// 	repo := infra.NewUserRepository(db)

// 	// 正常系
// 	testUser, err := repo.Update(context.Background(), *validUser)
// 	assert.NoError(t, err)
// 	assert.Equal(t, validUser, testUser)

// 	// 異常系
// 	_, err = repo.Update(context.Background(), *invalidUser)
// 	assert.Error(t, err)

// }

// func TestUserRepository_Delete(t *testing.T) {
// 	db, mock, err := sqlxmock.Newx()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// 正常系
// 	validUser := &entity.User{
// 		UID:            "uid",
// 		Username:       "username",
// 		Email:          "email",
// 		HashedPassword: "password",
// 		Text:           "text",
// 		Avater:         "avater",
// 		Header:         "header",
// 		CreatedAt:      "2020-01-01 00:00:00",
// 	}

// 	// モック
// 	mock.E

// 	// リポジトリの生成
// 	repo := infra.NewUserRepository(db)

// 正常系
// user, err := repo.Insert(context.Background(), *validUser)
// if err != nil {
// 	t.Fatalf("an error '%s' was not expected when insert user", err)
// }

// err = repo.DeleteTX(context.Background(), validUser.UID)
// assert.NoError(t, err)

// }
