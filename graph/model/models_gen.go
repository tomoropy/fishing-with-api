// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type User struct {
	UID       string `json:"uid"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Text      string `json:"text"`
	Avater    string `json:"avater"`
	Header    string `json:"header"`
	CreatedAt string `json:"createdAt"`
}

type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Text     string `json:"text"`
	Avater   string `json:"avater"`
	Header   string `json:"header"`
}
