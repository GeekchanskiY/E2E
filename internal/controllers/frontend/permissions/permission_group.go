package permissions

import (
	"context"
	"html/template"

	"go.uber.org/zap"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/templates"
)

func (c *controller) PermissionGroup(ctx context.Context, permissionGroupID int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.permission_group.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.PermissionGroupTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	permissionGroup, err := c.permissionGroupRepo.GetByID(ctx, permissionGroupID)
	if err != nil {
		return nil, nil, err
	}

	userPermissions, err := c.userPermissionRepo.GetForGroup(ctx, permissionGroup.ID)
	if err != nil {
		return nil, nil, err
	}

	data["group"] = permissionGroup
	data["permissions"] = userPermissions

	return html, data, nil
}
