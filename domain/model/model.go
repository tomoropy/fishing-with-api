package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	ProfileText    string `json:"profile_text"`
	ProfileImage   string `json:"profile_image"`
	HeaderImage    string `json:"header_image"`
}

type Invitation struct {
	gorm.Model
	UserID    int    `json:"user_id"`
	Commnet   string `json:"comment"`
	Place     string `json:"place"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Photo struct {
	gorm.Model
	InvitationID int    `json:"invitation_id"`
	Image        string `json:"image"`
}
