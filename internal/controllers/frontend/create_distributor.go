package frontend

import (
	"context"
	"errors"
	"html/template"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

func (c *Controller) CreateDistributor(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_distributor.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateDistributorTemplate)
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

	wallets, err := c.walletsRepo.GetByUsername(ctx, user)
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}
	data["wallets"] = wallets

	return html, data, nil
}

func (c *Controller) CreateDistributorForm(ctx context.Context, distributor models.Distributor) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_distributor.controller.form", zap.String("event", "got request"))

	err := distributor.Validate()
	if err != nil {
		return c.CreateDistributorFormError(ctx, err)
	}

	_, err = c.distributorsRepo.Create(ctx, &distributor)
	if err != nil {
		return c.CreateDistributorFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *Controller) CreateDistributorFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateDistributorTemplate)
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

	wallets, err := c.walletsRepo.GetByUsername(ctx, user)
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}
	data["wallets"] = wallets

	data["error"] = userErr.Error()

	return html, data, userErr
}
