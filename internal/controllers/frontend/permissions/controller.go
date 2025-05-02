package permissions

import (
	"context"
	"embed"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/models"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/templates"
)

type Controller interface {
	AddUser(ctx context.Context) (*template.Template, map[string]any, error)
	AddUserForm(ctx context.Context, username string, level string, permissionGroupID int64) (*template.Template, map[string]any, error)
	CreatePermissionGroup(ctx context.Context) (*template.Template, map[string]any, error)
	CreatePermissionGroupForm(ctx context.Context, permissionGroup *models.PermissionGroup) (*template.Template, map[string]any, error)
	List(ctx context.Context) (*template.Template, map[string]any, error)
	PermissionGroup(ctx context.Context, permissionGroupID int64) (*template.Template, map[string]any, error)
	DeleteUser(ctx context.Context, username string, permissionGroupID int64) error
}

type controller struct {
	logger *zap.Logger

	userRepo            users.Repository
	permissionGroupRepo permission_groups.Repository
	userPermissionRepo  user_permissions.Repository

	fs embed.FS
}

func New(
	logger *zap.Logger,

	userRepo users.Repository,
	permissionGroupRepo permission_groups.Repository,
	userPermissionRepo user_permissions.Repository,
) Controller {
	return &controller{
		logger: logger,

		userRepo:            userRepo,
		permissionGroupRepo: permissionGroupRepo,
		userPermissionRepo:  userPermissionRepo,

		fs: templates.Fs,
	}
}
