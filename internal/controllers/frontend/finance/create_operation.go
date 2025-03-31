package finance

import (
	"context"
	"errors"
	"fmt"
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
		operations   []*models.Operation
		distributors []*models.DistributorExtended

		errs []error
		err  error
	)

	operation.InitiatorId = ctx.Value("userId").(int64)

	if err = operation.Validate(); err != nil {
		return c.createOperationFormError(ctx, walletId, err)
	}

	if distributors, err = c.distributorsRepo.GetForWallet(ctx, walletId); err != nil {
		c.logger.Error("frontend.create_operation.controller.get_distributors", zap.Error(err))
		return c.createOperationFormError(ctx, walletId, err)
	}

	if len(distributors) != 0 {
		for _, d := range distributors {
			var (
				amountToWallet float64

				targetWallet, sourceWallet               models.Wallet
				targetCurrencyState, sourceCurrencyState *models.CurrencyState
				operationGroup                           *models.OperationGroup
			)

			amountToWallet = (operation.Amount / 100) * d.Percent
			operation.Amount -= amountToWallet

			if targetWallet, err = c.walletsRepo.Get(ctx, d.TargetWalletId); err != nil {
				c.logger.Error("frontend.create_operation.controller.get_wallet", zap.Error(err))

				return c.createOperationFormError(ctx, walletId, err)
			}

			if sourceWallet, err = c.walletsRepo.Get(ctx, d.SourceWalletId); err != nil {
				c.logger.Error("frontend.create_operation.controller.get_wallet", zap.Error(err))

				return c.createOperationFormError(ctx, walletId, err)
			}

			// Check wallet currency difference from internal database
			if sourceWallet.Currency != targetWallet.Currency {
				targetCurrencyState, sourceCurrencyState = new(models.CurrencyState), new(models.CurrencyState)

				if targetWallet.Currency != models.CurrencyUSD {
					if targetCurrencyState, err = c.currencyStatesRepo.GetBankCurrencyState(ctx, targetWallet.Currency, targetWallet.BankId); err != nil {
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
					if sourceCurrencyState, err = c.currencyStatesRepo.GetBankCurrencyState(ctx, sourceWallet.Currency, sourceWallet.BankId); err != nil {
						c.logger.Error("frontend.create_operation.controller.get_currency_states", zap.Error(err))

						return c.createOperationFormError(ctx, walletId, err)
					}
				} else {
					sourceCurrencyState = &models.CurrencyState{
						BuyUsd:  1,
						SellUsd: 1,
					}
				}

				if sourceWallet.Currency == models.CurrencyUSD && targetWallet.Currency != models.CurrencyUSD {
					if targetWallet.Currency != models.CurrencyUSD {
						if targetCurrencyState.SellUsd == 0 {
							return c.createOperationFormError(ctx, walletId, fmt.Errorf("cant get target currency SellUsd for %s", targetCurrencyState.CurrencyName))
						}

						amountToWallet = amountToWallet * targetCurrencyState.SellUsd
					}
				} else if targetWallet.Currency == models.CurrencyUSD && sourceWallet.Currency != models.CurrencyUSD {
					if sourceCurrencyState.BuyUsd == 0 {
						return c.createOperationFormError(ctx, walletId, fmt.Errorf("cant get target currency BuyUsd for %s", sourceCurrencyState.CurrencyName))
					}

					amountToWallet = amountToWallet * sourceCurrencyState.BuyUsd
				} else {
					if sourceCurrencyState.BuyUsd == 0 {
						return c.createOperationFormError(ctx, walletId, fmt.Errorf("cant get source currency BuyUsd for %s", sourceCurrencyState.CurrencyName))
					}

					if targetCurrencyState.SellUsd == 0 {
						return c.createOperationFormError(ctx, walletId, fmt.Errorf("cant get target currency SellUsd for %s", targetCurrencyState.CurrencyName))
					}

					amountToWallet = (amountToWallet * sourceCurrencyState.BuyUsd) * targetCurrencyState.SellUsd
				}
			}

			if operationGroup, err = c.operationGroupsRepo.GetOrCreateForWalletByName(ctx, d.TargetWalletId, "distributed"); err != nil {
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
