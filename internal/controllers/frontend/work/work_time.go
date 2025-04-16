package work

import (
	"context"
	"errors"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/config"
	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *controller) WorkTime(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.work_time.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.WorkTimeTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	user, ok := ctx.Value(config.UsernameContextKey).(string)
	if user == "" || !ok {
		err = errors.New("user is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	return html, data, nil
}

func (c *controller) WorkTimeForm(ctx context.Context, workID, userID int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.work_time.controller.form", zap.String("event", "got request"))

	if _, err := c.workRepo.StartWorkTime(ctx, workID); err != nil {
		return c.createWorkTimeFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *controller) createWorkTimeFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.WorkTimeTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	user, ok := ctx.Value(config.UsernameContextKey).(string)
	if user == "" || !ok {
		err = errors.New("user is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	data["error"] = userErr.Error()

	return html, data, userErr
}
