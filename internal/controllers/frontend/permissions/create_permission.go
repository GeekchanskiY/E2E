package permissions

import (
	"context"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/config"
	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"
)

func (c *controller) CreatePermissionGroup(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_permission_group.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreatePermissionGroupTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *controller) CreatePermissionGroupForm(ctx context.Context, permissionGroup *models.PermissionGroup) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_permission_group.controller.form", zap.String("event", "got request"))

	newPermissionGroup, err := c.permissionGroupRepo.Create(ctx, permissionGroup)
	if err != nil {
		return c.createPermissionGroupFormError(ctx, err)
	}

	userID := ctx.Value(config.UserIDContextKey).(int64)

	_, err = c.userPermissionRepo.Create(ctx, &models.UserPermission{
		PermissionGroupID: newPermissionGroup.ID,
		UserID:            userID,
		Level:             "owner",
	})
	if err != nil {
		return c.createPermissionGroupFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *controller) createPermissionGroupFormError(ctx context.Context, err error) (*template.Template, map[string]any, error) {
	html, generateErr := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreatePermissionGroupTemplate)
	if generateErr != nil {
		return nil, nil, generateErr
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)
	data["error"] = err

	return html, data, nil
}
