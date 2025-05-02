package finance

import (
	"context"
	"errors"
	"html/template"

	"finworker/internal/config"
	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/templates"

	"go.uber.org/zap"
)

func (c *controller) prepareDistributorData(ctx context.Context, data map[string]any) (map[string]any, error) {
	var err error

	user, ok := ctx.Value(config.UsernameContextKey).(string)
	if !ok {
		err = errors.New("user is empty")
		data["error"] = err.Error()

		return data, err
	}

	wallets, err := c.walletsRepo.GetByUsername(ctx, user)
	if err != nil {
		data["error"] = err.Error()

		return data, err
	}

	data["wallets"] = wallets

	return data, nil
}

func (c *controller) CreateDistributor(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_distributor.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateDistributorTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	data, err = c.prepareDistributorData(ctx, data)
	if err != nil {
		return html, data, err
	}

	return html, data, nil
}

func (c *controller) CreateDistributorForm(ctx context.Context, distributor *models.Distributor) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_distributor.controller.form", zap.String("event", "got request"))

	err := distributor.Validate()
	if err != nil {
		return c.createDistributorFormError(ctx, err)
	}

	_, err = c.distributorsRepo.Create(ctx, distributor)
	if err != nil {
		return c.createDistributorFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *controller) createDistributorFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateDistributorTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	data, err = c.prepareDistributorData(ctx, data)
	if err != nil {
		return html, data, err
	}

	data["error"] = userErr.Error()

	return html, data, userErr
}
