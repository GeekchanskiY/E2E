package permissions

import (
	"context"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/config"
	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *controller) List(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.list.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.ListPermissionGroupsTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	userPermissions, err := c.permissionGroupRepo.GetUserGroups(ctx, ctx.Value(config.UserIDContextKey).(int64))
	if err != nil {
		return nil, nil, err
	}

	data["permissions"] = userPermissions

	return html, data, nil
}
