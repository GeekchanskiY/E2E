package models

import "time"

type Wallet struct {
	Id                int       `db:"id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	PermissionGroupId int       `db:"permission_group_id"`
	CreatedAt         time.Time `db:"created_at"`
	Currency          string    `db:"currency"`
	IsSalary          bool      `db:"is_salary"`
}
