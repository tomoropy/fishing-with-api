package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	var user entity.User

	user.UID = uuid.New().String()
	user.Username = input.Username
	user.Email = input.Email
	user.HashedPassword = input.Password
	user.Avater = input.Avater
	user.Header = input.Header
	user.Text = input.Text
	user.CreatedAt = time.Now().Format(time.RFC3339)

	createdUser, err := r.MS.SaveUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	modelUser := r.P.User(createdUser)

	return modelUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	updatedUser, err := r.MS.SaveUser(ctx, &entity.User{
		UID:            input.UID,
		Username:       input.User.Username,
		Email:          input.User.Email,
		HashedPassword: input.User.Password,
		Avater:         input.User.Avater,
		Header:         input.User.Header,
		Text:           input.User.Text,
	})
	if err != nil {
		return nil, err
	}
	return r.P.User(updatedUser), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, uid string) (*model.ResponceInfo, error) {
	err := r.MS.DeleteUser(ctx, uid)
	if err != nil {
		return &model.ResponceInfo{
			Message: "削除に失敗しました",
			Status:  500,
		}, err
	}
	return &model.ResponceInfo{
		Message: "削除に成功しました",
		Status:  200,
	}, nil
}

// CreateTweet is the resolver for the createTweet field.
func (r *mutationResolver) CreateTweet(ctx context.Context, input model.TweetInput) (*model.Tweet, error) {
	createdTweet, err := r.MS.SaveTweet(ctx, &entity.Tweet{
		UID:       uuid.New().String(),
		UserUID:   input.UserID,
		Body:      input.Body,
		Image:     input.Image,
		CreatedAt: time.Now().Format("0000-00-00 00:00:00"),
	})
	if err != nil {
		return nil, err
	}
	return r.P.Tweet(createdTweet), nil
}

// UpdateTweet is the resolver for the updateTweet field.
func (r *mutationResolver) UpdateTweet(ctx context.Context, input model.UpdateTweetInput) (*model.Tweet, error) {
	updatedTweet, err := r.MS.SaveTweet(ctx, &entity.Tweet{
		UID:     input.UID,
		UserUID: input.Tweet.UserID,
		Body:    input.Tweet.Body,
		Image:   input.Tweet.Image,
	})
	if err != nil {
		return nil, err
	}
	return r.P.Tweet(updatedTweet), nil
}

// DeleteTweet is the resolver for the deleteTweet field.
func (r *mutationResolver) DeleteTweet(ctx context.Context, uid string) (*model.ResponceInfo, error) {
	err := r.MS.DeleteTweet(ctx, uid)
	if err != nil {
		return &model.ResponceInfo{
			Message: "削除に失敗しました",
			Status:  500,
		}, err
	}
	return &model.ResponceInfo{
		Message: "削除に成功しました",
		Status:  200,
	}, nil
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, email string, password string) (*model.User, error) {
	user, err := r.QS.Login(ctx, email, password)
	if err != nil {
		return nil, err
	}
	modelUser := r.P.User(user)
	return modelUser, nil
}

// AllUser is the resolver for the allUser field.
func (r *queryResolver) AllUser(ctx context.Context) ([]*model.User, error) {
	users, err := r.QS.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var modelUsers []*model.User
	for _, user := range users {
		modelUsers = append(modelUsers, r.P.User(&user))
	}

	return modelUsers, nil
}

// UserByUID is the resolver for the userByUid field.
func (r *queryResolver) UserByUID(ctx context.Context, uid string) (*model.User, error) {
	user, err := r.QS.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	modelUser := r.P.User(user)
	return modelUser, nil
}

// AllTweet is the resolver for the allTweet field.
func (r *queryResolver) AllTweet(ctx context.Context) ([]*model.Tweet, error) {
	tweets, err := r.QS.ListTweets(ctx)
	if err != nil {
		return nil, err
	}

	var modelTweets []*model.Tweet
	for _, tweet := range tweets {
		modelTweets = append(modelTweets, r.P.Tweet(&tweet))
	}
	return modelTweets, nil
}

// TweetsByUID is the resolver for the tweetsByUID field.
func (r *queryResolver) TweetsByUID(ctx context.Context, uid string) (*model.Tweet, error) {
	tweet, err := r.QS.GetTweet(ctx, uid)
	if err != nil {
		return nil, err
	}

	modelTweet := r.P.Tweet(tweet)
	return modelTweet, nil
}

// TweetByUserID is the resolver for the tweetByUserID field.
func (r *queryResolver) TweetByUserID(ctx context.Context, userid string) ([]*model.Tweet, error) {
	tweets, err := r.QS.GetTweetsByUserUID(ctx, userid)
	if err != nil {
		return nil, err
	}

	var modelTweets []*model.Tweet
	for _, tweet := range tweets {
		modelTweets = append(modelTweets, r.P.Tweet(&tweet))
	}
	return modelTweets, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
