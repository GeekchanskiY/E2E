package models

import "time"

type User struct {
	Id       int       `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password_hash"`
	Name     string    `db:"name"`
	Gender   string    `db:"gender"`
	Age      int       `db:"age"`
	Birthday time.Time `db:"birthday"`
}
