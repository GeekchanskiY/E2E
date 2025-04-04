package permissions

import (
	"context"
	"errors"
	"html/template"

	"go.uber.org/zap"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"
)

func (c *controller) AddUser(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.add_user_to_permission_group.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.AddUserToPermissionGroupTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *controller) AddUserForm(ctx context.Context, username string, level string, permissionGroupId int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.add_user_to_permission_group.controller.form", zap.String("event", "got request"))

	var (
		user           *models.User
		userPermission *models.UserPermission
		accessLevel    models.AccessLevel

		err error
	)

	if user, err = c.userRepo.GetByUsername(ctx, username); err != nil {
		return c.createPermissionGroupFormError(ctx, err)
	}

	accessLevel = models.AccessLevel(level)
	if !accessLevel.IsValid() {
		return c.createPermissionGroupFormError(ctx, errors.New("invalid access level"))
	}

	userPermission = new(models.UserPermission)
	userPermission.UserId = user.Id
	userPermission.PermissionGroupId = permissionGroupId
	userPermission.Level = accessLevel

	if _, err := c.userPermissionRepo.Create(ctx, userPermission); err != nil {
		return c.addUserFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *controller) addUserFormError(ctx context.Context, err error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.AddUserToPermissionGroupTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)
	data["error"] = err

	return html, data, nil
}
