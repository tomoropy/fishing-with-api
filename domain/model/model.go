package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username" gorm:"not null; unique"`
	Email          string `json:"email" gorm:"not null; unique"`
	HashedPassword string `json:"password" gorm:"not null; unique"`
	Text           string `json:"text"`
	Avater         string `json:"avater"`
	Header         string `json:"header"`
}

type Invitation struct {
	gorm.Model
	UserID  int    `json:"user_id" gorm:"not null"`
	Comment string `json:"comment" gorm:"not null"`
	Place   string `json:"place" gorm:"not null"`
}

// type Photo struct {
// 	gorm.Model
// 	InvitationID int    `json:"invitation_id"`
// 	Image        string `json:"image"`
// }
