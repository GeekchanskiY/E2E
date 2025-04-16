package models

import "time"

type User struct {
	ID       int64     `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password_hash" json:"password,omitempty"`
	Name     string    `db:"name" json:"name"`
	Gender   string    `db:"gender" json:"gender"`
	Birthday time.Time `db:"birthday" json:"birthday"`
}
