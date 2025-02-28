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

	// Wallet created initial wallet for salary
	Wallet *models.Wallet `json:"wallet,omitempty"`

	// OperationGroup created initial operation group for salary if salary data provided
	OperationGroup *models.OperationGroup `json:"operation_group,omitempty"`

	// Operation is a salary operation, which is monthly if specified
	Operation *models.Operation `json:"operation,omitempty"`
}
