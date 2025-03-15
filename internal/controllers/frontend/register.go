package frontend

import (
	"context"
	"errors"
	"html/template"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

func (c *Controller) Register(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *Controller) RegisterForm(ctx context.Context, username, password, repeatPassword, gender, birthday, bank, salary, currency, payday string) (*template.Template, map[string]any, string, string, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, "", "", err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	data["error"] = "form not ready"
	return html, data, "", "", errors.New("form not ready")
}
