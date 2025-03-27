package finance

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

func (c *controller) CreateOperation(ctx context.Context, walletId int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateOperationTemplate)
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

	wallet, err := c.walletsRepo.Get(ctx, walletId)
	if err != nil {
		data["error"] = err.Error()
		return html, data, err
	}
	data["wallet"] = wallet

	return html, data, nil
}

func (c *controller) CreateOperationForm(ctx context.Context, operation models.Operation, walletId int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation.controller.form", zap.String("event", "got request"))

	var (
		operations []*models.Operation
		errs       []error
	)

	operation.InitiatorId = ctx.Value("userId").(int64)

	err := operation.Validate()
	if err != nil {
		return c.createOperationFormError(ctx, walletId, err)
	}

	distributors, err := c.distributorsRepo.GetForWallet(ctx, walletId)
	if err != nil {
		c.logger.Error("frontend.create_operation.controller.get_distributors", zap.Error(err))
		return c.createOperationFormError(ctx, walletId, err)
	}

	summaryAmountMinus := float64(0)

	if len(distributors) != 0 {
		for _, d := range distributors {
			amountToWallet := (operation.Amount / 100) * d.Percent
			summaryAmountMinus += amountToWallet

			targetWallet, err := c.walletsRepo.Get(ctx, d.TargetWalletId)
			if err != nil {
				c.logger.Error("frontend.create_operation.controller.get_wallet", zap.Error(err))

				return c.createOperationFormError(ctx, walletId, err)
			}

			sourceWallet, err := c.walletsRepo.Get(ctx, d.SourceWalletId)
			if err != nil {
				c.logger.Error("frontend.create_operation.controller.get_wallet", zap.Error(err))

				return c.createOperationFormError(ctx, walletId, err)
			}

			// Check wallet currency difference from internal database
			if sourceWallet.Currency != targetWallet.Currency {
				targetCurrencyState, sourceCurrencyState := new(models.CurrencyState), new(models.CurrencyState)

				if targetWallet.Currency != models.CurrencyUSD {
					targetCurrencyState, err = c.currencyStatesRepo.GetBankCurrencyState(ctx, targetWallet.Currency, targetWallet.BankId)
					if err != nil {
						c.logger.Error("frontend.create_operation.controller.get_currency_states", zap.Error(err))

						return c.createOperationFormError(ctx, walletId, err)
					}
				} else {
					targetCurrencyState = &models.CurrencyState{
						BuyUsd:  1,
						SellUsd: 1,
					}
				}

				if sourceWallet.Currency != models.CurrencyUSD {
					sourceCurrencyState, err = c.currencyStatesRepo.GetBankCurrencyState(ctx, sourceWallet.Currency, sourceWallet.BankId)
					if err != nil {
						c.logger.Error("frontend.create_operation.controller.get_currency_states", zap.Error(err))

						return c.createOperationFormError(ctx, walletId, err)
					}
				} else {
					sourceCurrencyState = &models.CurrencyState{
						BuyUsd:  1,
						SellUsd: 1,
					}
				}

				amountToWallet = (amountToWallet * sourceCurrencyState.SellUsd) * targetCurrencyState.BuyUsd
			}

			operationGroup, err := c.operationGroupsRepo.GetOrCreateForWalletByName(ctx, d.TargetWalletId, "distributed")
			if err != nil {
				c.logger.Error("frontend.create_operation.controller.get_or_create_operation_groups", zap.Error(err))

				return c.createOperationFormError(ctx, walletId, err)
			}
			operations = append(operations, &models.Operation{
				OperationGroupId: operationGroup.Id,
				Time:             time.Now(),
				IsMonthly:        false,
				IsConfirmed:      true,
				Amount:           amountToWallet,
				InitiatorId:      ctx.Value("userId").(int64),
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
		return c.createOperationFormError(ctx, walletId, err)
	}

	operation.Amount -= summaryAmountMinus
	_, err = c.operationsRepo.Create(ctx, &operation)
	if err != nil {
		return c.createOperationFormError(ctx, walletId, err)
	}

	return nil, nil, nil
}

func (c *controller) createOperationFormError(ctx context.Context, walletId int64, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateOperationTemplate)
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

	wallet, err := c.walletsRepo.Get(ctx, walletId)
	if err != nil {
		data["error"] = err.Error()
		return html, data, err
	}
	data["wallet"] = wallet

	data["error"] = userErr.Error()

	return html, data, userErr
}
