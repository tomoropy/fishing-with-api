package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	Text           string `json:"text"`
	Avater         string `json:"avater"`
	Header         string `json:"header"`
}

type Invitation struct {
	gorm.Model
	UserID  int    `json:"user_id"`
	Comment string `json:"comment"`
	Place   string `json:"place"`
}

// type Photo struct {
// 	gorm.Model
// 	InvitationID int    `json:"invitation_id"`
// 	Image        string `json:"image"`
// }
