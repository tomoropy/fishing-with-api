package entity

type User struct {
	UID            string `json:"uid"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	Text           string `json:"text"`
	Avater         string `json:"avater"`
	Header         string `json:"header"`
	CreatedAt      string `json:"created_at"`
}

type Invitation struct {
	UID       string `json:"uid"`
	UserUID   string `json:"user_uid"`
	Comment   string `json:"comment"`
	Place     string `json:"place"`
	CreatedAt string `json:"created_at"`
}
