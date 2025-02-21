package models

import "time"

type AccessLevel string

const (
	AccessLevelOwner AccessLevel = "owner"
	AccessLevelFull  AccessLevel = "full"
	AccessLevelRead  AccessLevel = "read"
)

type UserPermission struct {
	Id                int         `db:"id"  json:"id"`
	PermissionGroupId int         `db:"permission_group_id" json:"permission_group_id"`
	UserId            int         `db:"user_id"  json:"user_id"`
	Level             AccessLevel `db:"level"   json:"level"`
	CreatedAt         time.Time   `db:"created_at" json:"created_at"`
}
