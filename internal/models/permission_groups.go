package models

import "time"

type PermissionGroup struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	Level     string    `db:"level"`
}
