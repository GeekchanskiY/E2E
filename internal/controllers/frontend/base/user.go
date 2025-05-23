package base

import (
	"context"
	"html/template"

	"finworker/internal/controllers/frontend/utils"
	"finworker/templates"

	"go.uber.org/zap"
)

func (c *controller) User(ctx context.Context, username string) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.user.controller", zap.String("event", "got request"))

	data := utils.BuildDefaultDataMapFromContext(ctx)

	data["is_me"] = data["username"] == username

	userData, err := c.userRepo.GetByUsername(ctx, username)
	if err != nil {
		c.logger.Error("frontend.user.controller", zap.Error(err))
		return nil, nil, err
	}

	data["userID"] = userData.ID
	data["userGender"] = userData.Gender
	data["userBirthday"] = userData.Birthday
	data["avatar"] = userData.Avatar

	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.UserTemplate)
	if err != nil {
		return nil, nil, err
	}

	return html, data, nil
}
