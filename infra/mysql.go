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

func (ur *userRepository) SelectAll(ctx context.Context) ([]entity.User, error) {

	var users []entity.User

	findAllUserSql := "SELECT * FROM users"

	err := ur.db.Select(&users, findAllUserSql)
	if err != nil {
		log.Error("failed to select all user" + err.Error())
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) SelectByUID(ctx context.Context, uid string) (*entity.User, error) {

	var user entity.User

	findUserByUIDSql := "SELECT * FROM users WHERE uid = ?"

	err := ur.db.Get(&user, findUserByUIDSql, uid)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) SelectByEmail(ctx context.Context, email string) (*entity.User, error) {

	var user entity.User

	findUserByEmailSql := "SELECT * FROM users WHERE email = ?"

	err := ur.db.Get(&user, findUserByEmailSql, email)
	if err != nil {
		log.Error("failed to select user by email" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Insert(ctx context.Context, user entity.User) (*entity.User, error) {

	insertUserSql := "INSERT INTO users (uid, username, email, hashed_password, text, avater, header, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := ur.db.Exec(insertUserSql, user.UID, user.Username, user.Email, user.HashedPassword, user.Text, user.Avater, user.Header, user.CreatedAt)
	if err != nil {
		log.Error("failed to insert user" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Update(ctx context.Context, user entity.User) (*entity.User, error) {

	updateUserSql := "UPDATE users SET username = ?, email = ?, hashed_password = ?, text = ?, avater = ?, header = ? WHERE uid = ?"

	_, err := ur.db.Exec(updateUserSql, user.Username, user.Email, user.HashedPassword, user.Text, user.Avater, user.Header, user.UID)
	if err != nil {
		log.Error("failed to update user" + err.Error())
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) DeleteTX(ctx context.Context, uid string) error {

	tx, err := ur.db.Beginx()
	if err != nil {
		log.Error("failed to begin transaction" + err.Error())
		return err
	}

	deleteUserSql := "DELETE FROM users WHERE uid = ?"
	deleteTweetSql := "DELETE FROM tweets WHERE user_uid = ?"

	_, err = tx.Exec(deleteTweetSql, uid)
	if err != nil {
		log.Error("failed to delete tweet" + err.Error())
		return err
	}

	_, err = tx.Exec(deleteUserSql, uid)
	if err != nil {
		log.Error("failed to delete user" + err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Error("failed to commit transaction" + err.Error())
		return err
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------------------------
// tweet repository
type tweetRepository struct {
	db *sqlx.DB
}

func NewTweetRepository(db *sqlx.DB) repository.TweetRepository {
	return &tweetRepository{
		db: db,
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------------

func (tr *tweetRepository) SelectAll(ctx context.Context) ([]entity.Tweet, error) {

	var tweets []entity.Tweet

	findAllTweetSql := "SELECT * FROM tweets"

	err := tr.db.Select(&tweets, findAllTweetSql)
	if err != nil {
		log.Error("failed to select all tweet" + err.Error())
		return nil, err
	}

	return tweets, nil
}

func (tr *tweetRepository) SelectByUID(ctx context.Context, uid string) (*entity.Tweet, error) {

	var tweet entity.Tweet

	findTweetByUIDSql := "SELECT * FROM tweets WHERE uid = ?"

	err := tr.db.Get(&tweet, findTweetByUIDSql, uid)
	if err != nil {
		log.Error("failed to select tweet by uid " + err.Error())
		return nil, err
	}

	return &tweet, nil
}

func (tr *tweetRepository) SelectByUserUID(ctx context.Context, userUID string) ([]entity.Tweet, error) {

	var tweets []entity.Tweet

	findTweetByUserUIDSql := "SELECT * FROM tweets WHERE user_uid = ?"

	err := tr.db.Select(&tweets, findTweetByUserUIDSql, userUID)
	if err != nil {
		log.Error("failed to select tweet by user uid " + err.Error())
		return nil, err
	}

	return tweets, nil
}

func (tr *tweetRepository) Insert(ctx context.Context, tweet entity.Tweet) (*entity.Tweet, error) {

	insertTweetSql := "INSERT INTO tweets (uid, user_uid, body, image, created_at) VALUES (?, ?, ?, ?, ?)"

	_, err := tr.db.Exec(insertTweetSql, tweet.UID, tweet.UserUID, tweet.Body, tweet.Image, tweet.CreatedAt)
	if err != nil {
		log.Error("failed to insert tweet" + err.Error())
		return nil, err
	}

	return &tweet, nil
}

func (tr *tweetRepository) Update(ctx context.Context, tweet entity.Tweet) (*entity.Tweet, error) {

	updateTweetSql := "UPDATE tweets SET body = ?, image = ? WHERE uid = ?"

	_, err := tr.db.Exec(updateTweetSql, tweet.Body, tweet.Image, tweet.UID)
	if err != nil {
		log.Error("failed to update tweet" + err.Error())
		return nil, err
	}

	return &tweet, nil
}

func (tr *tweetRepository) Delete(ctx context.Context, uid string) error {

	deleteTweetSql := "DELETE FROM tweets WHERE uid = ?"

	_, err := tr.db.Exec(deleteTweetSql, uid)
	if err != nil {
		log.Error("failed to delete tweet" + err.Error())
		return err
	}

	return nil
}
