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

func (c *Controller) CreateOperationGroup(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation_group.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateOperationGroupTemplate)
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

func (c *Controller) CreateOperationGroupForm(ctx context.Context, operationGroup models.OperationGroup) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation_group.controller.form", zap.String("event", "got request"))

	err := operationGroup.Validate()
	if err != nil {
		return c.CreateOperationGroupFormError(ctx, err)
	}

	_, err = c.operationGroupsRepo.Create(ctx, &operationGroup)
	if err != nil {
		return c.CreateOperationGroupFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *Controller) CreateOperationGroupFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateOperationGroupTemplate)
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
