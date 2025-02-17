package models

type UserPermission struct {
	Id                int `db:"id"`
	PermissionGroupId int `db:"permission_group"`
	UserId            int `db:"user_id"`
}
