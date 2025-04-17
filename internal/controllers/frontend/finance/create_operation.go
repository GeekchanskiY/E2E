package finance

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"time"

	"finworker/internal/config"
	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

func (c *controller) prepareOperationData(ctx context.Context, walletID int64, data map[string]any) (map[string]any, error) {
	operationGroups, err := c.operationGroupsRepo.GetByWallet(ctx, walletID)
	if err != nil {
		data["error"] = err.Error()
		return data, err
	}

	data["operation_groups"] = operationGroups

	wallet, err := c.walletsRepo.Get(ctx, walletID)
	if err != nil {
		data["error"] = err.Error()
		return data, err
	}
	data["wallet"] = wallet

	return data, err
}

func (c *controller) CreateOperation(ctx context.Context, walletID int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateOperationTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	if data, err = c.prepareOperationData(ctx, walletID, data); err != nil {
		return html, data, err
	}

	return html, data, nil
}

func (c *controller) CreateOperationForm(ctx context.Context, operation *models.Operation, walletID int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_operation.controller.form", zap.String("event", "got request"))

	var (
		operations   []*models.Operation
		distributors []*models.DistributorExtended

		errs []error
		err  error
	)

	operation.InitiatorID = ctx.Value(config.UserIDContextKey).(int64)

	if err = operation.Validate(); err != nil {
		return c.createOperationFormError(ctx, walletID, err)
	}

	if distributors, err = c.distributorsRepo.GetForWallet(ctx, walletID); err != nil {
		c.logger.Error("frontend.create_operation.controller.get_distributors", zap.Error(err))
		return c.createOperationFormError(ctx, walletID, err)
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

			if targetWallet, err = c.walletsRepo.Get(ctx, d.TargetWalletID); err != nil {
				c.logger.Error("frontend.create_operation.controller.get_wallet", zap.Error(err))

				return c.createOperationFormError(ctx, walletID, err)
			}

			if sourceWallet, err = c.walletsRepo.Get(ctx, d.SourceWalletID); err != nil {
				c.logger.Error("frontend.create_operation.controller.get_wallet", zap.Error(err))

				return c.createOperationFormError(ctx, walletID, err)
			}

			// Check wallet currency difference from internal database
			if sourceWallet.Currency != targetWallet.Currency {
				targetCurrencyState, sourceCurrencyState = new(models.CurrencyState), new(models.CurrencyState)

				if targetWallet.Currency != models.CurrencyUSD {
					if targetCurrencyState, err = c.currencyStatesRepo.GetBankCurrencyState(ctx, targetWallet.Currency, targetWallet.BankID); err != nil {
						c.logger.Error("frontend.create_operation.controller.get_currency_states", zap.Error(err))

						return c.createOperationFormError(ctx, walletID, err)
					}
				} else {
					targetCurrencyState = &models.CurrencyState{
						BuyUsd:  1,
						SellUsd: 1,
					}
				}

				if sourceWallet.Currency != models.CurrencyUSD {
					if sourceCurrencyState, err = c.currencyStatesRepo.GetBankCurrencyState(ctx, sourceWallet.Currency, sourceWallet.BankID); err != nil {
						c.logger.Error("frontend.create_operation.controller.get_currency_states", zap.Error(err))

						return c.createOperationFormError(ctx, walletID, err)
					}
				} else {
					sourceCurrencyState = &models.CurrencyState{
						BuyUsd:  1,
						SellUsd: 1,
					}
				}

				switch sourceWallet.Currency {
				case models.CurrencyUSD:
					if targetWallet.Currency != models.CurrencyUSD {
						if targetCurrencyState.SellUsd == 0 {
							return c.createOperationFormError(ctx, walletID, fmt.Errorf("cant get target currency SellUsd for %s", targetCurrencyState.CurrencyName))
						}

						amountToWallet *= targetCurrencyState.SellUsd
					}
				default:
					if targetWallet.Currency == models.CurrencyUSD {
						if sourceCurrencyState.BuyUsd == 0 {
							return c.createOperationFormError(ctx, walletID, fmt.Errorf("cant get target currency BuyUsd for %s", sourceCurrencyState.CurrencyName))
						}

						amountToWallet *= sourceCurrencyState.BuyUsd
					}
				}

				if sourceWallet.Currency != models.CurrencyUSD && targetWallet.Currency != models.CurrencyUSD {
					if sourceCurrencyState.BuyUsd == 0 {
						return c.createOperationFormError(ctx, walletID, fmt.Errorf("cant get source currency BuyUsd for %s", sourceCurrencyState.CurrencyName))
					}

					if targetCurrencyState.SellUsd == 0 {
						return c.createOperationFormError(ctx, walletID, fmt.Errorf("cant get target currency SellUsd for %s", targetCurrencyState.CurrencyName))
					}

					amountToWallet = (amountToWallet * sourceCurrencyState.BuyUsd) * targetCurrencyState.SellUsd
				}
			}

			if operationGroup, err = c.operationGroupsRepo.GetOrCreateForWalletByName(ctx, d.TargetWalletID, "distributed"); err != nil {
				c.logger.Error("frontend.create_operation.controller.get_or_create_operation_groups", zap.Error(err))

				return c.createOperationFormError(ctx, walletID, err)
			}

			operations = append(operations, &models.Operation{
				OperationGroupID: operationGroup.ID,
				Time:             time.Now(),
				IsMonthly:        false,
				IsConfirmed:      true,
				Amount:           amountToWallet,
				InitiatorID:      ctx.Value(config.UserIDContextKey).(int64),
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
		return c.createOperationFormError(ctx, walletID, err)
	}

	_, err = c.operationsRepo.Create(ctx, operation)
	if err != nil {
		return c.createOperationFormError(ctx, walletID, err)
	}

	return nil, nil, nil
}

func (c *controller) createOperationFormError(ctx context.Context, walletID int64, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateOperationTemplate)
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

	if data, err = c.prepareOperationData(ctx, walletID, data); err != nil {
		return html, data, err
	}

	data["error"] = userErr.Error()

	return html, data, userErr
}
