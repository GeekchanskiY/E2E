package models

import "time"

type User struct {
	Id       int       `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password_hash" json:"password,omitempty"`
	Name     string    `db:"name" json:"name"`
	Gender   string    `db:"gender" json:"gender"`
	Birthday time.Time `db:"birthday" json:"birthday"`
}
