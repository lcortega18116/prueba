package models

type User struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	FullName string `db:"full_name" json:"full_name"`
}
