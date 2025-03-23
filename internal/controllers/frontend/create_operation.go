package frontend

import (
	"context"
	"errors"
	"html/template"
	"time"

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

	var (
		operations []*models.Operation
		errs       []error
	)

	operation.InitiatorId = int(ctx.Value("userId").(int64))

	err := operation.Validate()
	if err != nil {
		return c.CreateOperationFormError(ctx, walletId, err)
	}

	distributors, err := c.distributorsRepo.GetForWallet(ctx, walletId)
	if err != nil {
		c.logger.Error("frontend.create_operation.controller.get_distributors", zap.Error(err))
		return c.CreateOperationFormError(ctx, walletId, err)
	}

	if len(distributors) != 0 {
		for _, d := range distributors {
			percentToWallet := operation.Amount / d.Percent
			print(percentToWallet)

			// TODO: fill distributed operations
			operations = append(operations, &models.Operation{
				OperationGroupId: 0,
				IsConsumption:    false,
				Time:             time.Time{},
				IsMonthly:        false,
				IsConfirmed:      false,
				Amount:           0,
				InitiatorId:      0,
			})
		}
	}

	for _, o := range operations {
		_, err = c.operationsRepo.Create(ctx, o)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if err = errors.Join(errs...); err != nil {
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
