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

func (c *Controller) CreateOperation(ctx context.Context, walletId int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateOperationTemplate)
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

	operationGroups, err := c.operationGroupsRepo.GetByWallet(ctx, walletId)
	if err != nil {
		data["error"] = err.Error()
		return html, data, err
	}
	data["operation_groups"] = operationGroups

	wallet, err := c.walletsRepo.Get(ctx, int(walletId))
	if err != nil {
		data["error"] = err.Error()
		return html, data, err
	}
	data["wallet"] = wallet

	return html, data, nil
}

func (c *Controller) CreateOperationForm(ctx context.Context, operation models.Operation, walletId int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation.controller.form", zap.String("event", "got request"))

	operation.InitiatorId = int(ctx.Value("userId").(int64))

	err := operation.Validate()
	if err != nil {
		return c.CreateOperationFormError(ctx, walletId, err)
	}

	_, err = c.operationsRepo.Create(ctx, &operation)
	if err != nil {
		return c.CreateOperationFormError(ctx, walletId, err)
	}

	return nil, nil, nil
}

func (c *Controller) CreateOperationFormError(ctx context.Context, walletId int64, userErr error) (*template.Template, map[string]any, error) {
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateOperationTemplate)
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

	operationGroups, err := c.operationGroupsRepo.GetByWallet(ctx, walletId)
	if err != nil {
		data["error"] = err.Error()
		return html, data, err
	}
	data["operation_groups"] = operationGroups

	wallet, err := c.walletsRepo.Get(ctx, int(walletId))
	if err != nil {
		data["error"] = err.Error()
		return html, data, err
	}
	data["wallet"] = wallet

	data["error"] = userErr.Error()

	return html, data, userErr
}
