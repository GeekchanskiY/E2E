package models

import "time"

type AccessLevel string

const (
	AccessLevelOwner AccessLevel = "owner"
	AccessLevelFull  AccessLevel = "full"
	AccessLevelRead  AccessLevel = "read"
)

func (l AccessLevel) IsValid() bool {
	if l != AccessLevelFull && l != AccessLevelRead && l != AccessLevelOwner {
		return false
	}

	return true
}

type UserPermission struct {
	ID                int64       `db:"id"  json:"id"`
	PermissionGroupID int64       `db:"permission_group_id" json:"permission_group_id"`
	UserID            int64       `db:"user_id"  json:"user_id"`
	Level             AccessLevel `db:"level"   json:"level"`
	CreatedAt         time.Time   `db:"created_at" json:"created_at"`
}

type UserPermissionExtended struct {
	ID                int64       `db:"id"  json:"id"`
	PermissionGroupID int64       `db:"permission_group_id" json:"permission_group_id"`
	UserID            int64       `db:"user_id"  json:"user_id"`
	Username          string      `db:"username"  json:"username"`
	Level             AccessLevel `db:"level"   json:"level"`
	CreatedAt         time.Time   `db:"created_at" json:"created_at"`
}
