package work

import (
	"context"
	"errors"
	"html/template"

	"go.uber.org/zap"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"
)

func (c *controller) CreateWork(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.work_time.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateWorkTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	user, ok := ctx.Value("user").(string)
	if user == "" || !ok {
		err = errors.New("user is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	return html, data, nil
}

func (c *controller) CreateWorkForm(ctx context.Context, work *models.UserWork) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.work_time.controller.form", zap.String("event", "got request"))

	if _, err := c.workRepo.CreateUserWork(ctx, work); err != nil {
		return c.createWorkFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *controller) createWorkFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateWorkTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	user, ok := ctx.Value("user").(string)
	if user == "" || !ok {
		err = errors.New("user is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	data["error"] = userErr.Error()

	return html, data, userErr
}
