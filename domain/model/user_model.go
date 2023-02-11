package model

type User struct {
	ID       int    `db:"id"`
	Username     string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Age      int    `db:"age"`
}
