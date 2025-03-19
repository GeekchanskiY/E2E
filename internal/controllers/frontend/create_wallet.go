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

func (c *Controller) CreateWallet(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_wallet.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateWalletTemplate)
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

	userId := ctx.Value("userId").(int64)
	if userId == 0 {
		err = errors.New("userId is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	userPermissionGroups, err := c.permissionGroupsRepo.GetUserEditGroups(ctx, userId)
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}

	data["userPermissionGroups"] = userPermissionGroups

	return html, data, nil
}

func (c *Controller) CreateWalletForm(ctx context.Context, walletData models.WalletExtended) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_wallet.controller.form", zap.String("event", "got request"))

	err := walletData.Validate()
	if err != nil {
		return c.CreateWalletFormError(ctx, err)
	}

	bank, err := c.banksRepo.GetByName(ctx, walletData.BankName)
	if err != nil {
		return c.CreateWalletFormError(ctx, err)
	}

	permissionGroup, err := c.permissionGroupsRepo.GetByName(ctx, walletData.Permission)
	if err != nil {
		return c.CreateWalletFormError(ctx, err)
	}

	_, err = c.walletsRepo.Create(ctx, &models.Wallet{
		Name:              walletData.Name,
		Description:       walletData.Description,
		PermissionGroupId: permissionGroup.Id,
		Currency:          walletData.Currency,
		IsSalary:          walletData.IsSalary,
		BankId:            bank.Id,
	})
	if err != nil {
		return c.CreateWalletFormError(ctx, err)
	}

	return nil, nil, nil
}

func (c *Controller) CreateWalletFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, error) {
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateWalletTemplate)
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

	userId := ctx.Value("userId").(int64)
	if userId == 0 {
		err = errors.New("userId is empty")
		data["error"] = err.Error()

		return html, data, err
	}

	userPermissionGroups, err := c.permissionGroupsRepo.GetUserEditGroups(ctx, userId)
	if err != nil {
		data["error"] = err.Error()

		return html, data, err
	}

	data["userPermissionGroups"] = userPermissionGroups

	data["error"] = userErr.Error()

	return html, data, userErr
}
