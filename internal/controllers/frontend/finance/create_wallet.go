package finance

import (
	"context"
	"errors"
	"html/template"

	"finworker/internal/config"
	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

func (c *controller) CreateWallet(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_wallet.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateWalletTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	banks, err := c.banksRepo.GetAll()
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}

	data["banks"] = banks

	userID, ok := ctx.Value(config.UserIDContextKey).(int64)
	if userID == 0 || !ok {
		err = errors.New("userId is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	userPermissionGroups, err := c.permissionGroupsRepo.GetUserEditGroups(ctx, userID)
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}

	data["userPermissionGroups"] = userPermissionGroups

	return html, data, nil
}

func (c *controller) CreateWalletForm(ctx context.Context, walletData *models.WalletExtended) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_wallet.controller.form", zap.String("event", "got request"))

	err := walletData.Validate()
	if err != nil {
		return c.createWalletFormError(ctx, err)
	}

	bank, err := c.banksRepo.GetByName(ctx, walletData.BankName)
	if err != nil {
		return c.createWalletFormError(ctx, err)
	}

	permissionGroup, err := c.permissionGroupsRepo.GetByName(ctx, walletData.Permission)
	if err != nil {
		return c.createWalletFormError(ctx, err)
	}

	_, err = c.walletsRepo.Create(ctx, &models.Wallet{
		Name:              walletData.Name,
		Description:       walletData.Description,
		PermissionGroupID: permissionGroup.ID,
		Currency:          walletData.Currency,
		IsSalary:          walletData.IsSalary,
		BankID:            bank.ID,
	})
	if err != nil {
		return c.createWalletFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *controller) createWalletFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.CreateWalletTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	banks, err := c.banksRepo.GetAll()
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}

	data["banks"] = banks

	userID, ok := ctx.Value(config.UserIDContextKey).(int64)
	if userID == 0 || !ok {
		err = errors.New("userId is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	userPermissionGroups, err := c.permissionGroupsRepo.GetUserEditGroups(ctx, userID)
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}

	data["userPermissionGroups"] = userPermissionGroups

	data["error"] = userErr.Error()

	return html, data, userErr
}
