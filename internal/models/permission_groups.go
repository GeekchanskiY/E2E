package models

import "time"

type PermissionGroup struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name"  json:"name"`
	CreatedAt time.Time `db:"created_at"  json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type PermissionGroupWithRole struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name"  json:"name"`
	CreatedAt  time.Time `db:"created_at"  json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Level      string    `db:"level"  json:"level"`
	UsersCount int64     `db:"users_count"  json:"users_count"`
}
