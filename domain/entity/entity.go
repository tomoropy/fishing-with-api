package entity

type User struct {
	UID            string `json:"uid"        db:"uid"`
	Username       string `json:"username"   db:"username"`
	Email          string `json:"email"      db:"email"`
	HashedPassword string `json:"password"   db:"hashed_password"`
	Text           string `json:"text"       db:"text"`
	Avater         string `json:"avater"     db:"avater"`
	Header         string `json:"header"     db:"header"`
	CreatedAt      string `json:"created_at" db:"created_at"`
}

type Tweet struct {
	UID       string `json:"uid"        db:"uid"`
	UserUID   string `json:"user_uid"   db:"user_uid"`
	Body      string `json:"body"       db:"body"`
	Image     string `json:"image"      db:"image"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
