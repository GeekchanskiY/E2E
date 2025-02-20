package users

import (
	"finworker/internal/models"
)

type RegisterResponse struct {
	// User new user data
	User *models.User `json:"user,omitempty"`

	// PermissionGroup that contains user as owner
	PermissionGroup *models.PermissionGroup `json:"permission_group,omitempty"`

	// UserPermission confirms that user is owner of the permission group
	UserPermission *models.UserPermission `json:"user_permission,omitempty"`
}
