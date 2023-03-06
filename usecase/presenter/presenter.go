package presenter

import (
	"github.com/tomoropy/fishing-with-api/domain/entity"
	"github.com/tomoropy/fishing-with-api/graph/model"
)

type presenter struct{}

type Presenter interface {
	User(*entity.User) *model.User
	Tweet(*entity.Tweet) *model.Tweet
}

func NewPresenter() Presenter {
	return &presenter{}
}

func (p *presenter) User(user *entity.User) *model.User {
	return &model.User{
		UID:       user.UID,
		Username:  user.Username,
		Password:  "", // password is not returned
		Email:     user.Email,
		Text:      user.Text,
		Avater:    user.Avater,
		Header:    user.Header,
		CreatedAt: user.CreatedAt,
	}
}

func (p *presenter) Tweet(tweet *entity.Tweet) *model.Tweet {
	return &model.Tweet{
		UID:       tweet.UID,
		UserID:    tweet.UserUID,
		Body:      tweet.Body,
		Image:     tweet.Image,
		CreatedAt: tweet.CreatedAt,
	}
}
