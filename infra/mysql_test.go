package infra_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"

	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/infra"
)

func TestUserRepository_Insert(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 正常系
	validUser := &entity.User{
		UID:            "uid",
		Username:       "username",
		Email:          "email",
		HashedPassword: "password",
		Text:           "text",
		Avater:         "avater",
		Header:         "header",
		CreatedAt:      "2020-01-01 00:00:00",
	}

	// 異常系(すでに登録されているユーザー)
	duplicateUser := &entity.User{
		UID:            "uid",
		Username:       "username",
		Email:          "email",
		HashedPassword: "password",
		Text:           "text",
		Avater:         "avater",
		Header:         "header",
		CreatedAt:      "2020-01-01 00:00:00",
	}

	// モック
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (uid, username, email, hashed_password, text, avater, header, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")).
		WithArgs("uid", "username", "email", "password", "text", "avater", "header", "2020-01-01 00:00:00").
		WillReturnResult(sqlxmock.NewResult(1, 1))

	// リポジトリの生成
	repo := infra.NewUserRepository(db)

	// 正常系
	testUser, err := repo.Insert(context.Background(), *validUser)
	assert.NoError(t, err)
	assert.Equal(t, validUser, testUser)

	// 異常系
	_, err = repo.Insert(context.Background(), *duplicateUser)
	assert.Error(t, err)

}

func TestUserRepository_Update(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// 正常系
	validUser := &entity.User{
		UID:            "uid",
		Username:       "username",
		Email:          "email",
		HashedPassword: "password",
		Text:           "text",
		Avater:         "avater",
		Header:         "header",
		CreatedAt:      "2020-01-01 00:00:00",
	}

	// 異常系(存在しないユーザー)
	invalidUser := &entity.User{
		UID:            "invalid_uid",
		Username:       "username",
		Email:          "email",
		HashedPassword: "password",
		Text:           "text",
		Avater:         "avater",
		Header:         "header",
		CreatedAt:      "2020-01-01 00:00:00",
	}

	// モック
	mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET username = ?, email = ?, hashed_password = ?, text = ?, avater = ?, header = ? WHERE uid = ?")).
		WithArgs("username", "email", "password", "text", "avater", "header", "uid").
		WillReturnResult(sqlxmock.NewResult(1, 1))

	// リポジトリの生成
	repo := infra.NewUserRepository(db)

	// 正常系
	testUser, err := repo.Update(context.Background(), *validUser)
	assert.NoError(t, err)
	assert.Equal(t, validUser, testUser)

	// 異常系
	_, err = repo.Update(context.Background(), *invalidUser)
	assert.Error(t, err)

}

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
