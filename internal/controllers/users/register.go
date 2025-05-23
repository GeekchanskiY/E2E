package users

import (
	"context"

	"finworker/internal/models"
	requests "finworker/internal/models/requests/users"
	responses "finworker/internal/models/responses/users"
	"finworker/internal/utils"
)

func (c *Controller) RegisterUser(ctx context.Context, req requests.RegisterRequest) (resp *responses.RegisterResponse, err error) {
	c.logger.Info("UserController.RegisterUser")

	bank, err := c.bankRepo.GetByName(ctx, req.PreferredBankName)
	if err != nil {
		return nil, err
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username: req.Username,
		Password: password,
		Name:     req.Name,
		Gender:   req.Gender,
		Birthday: req.Birthday,
	}

	newUser, err = c.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	// sets password to empty to avoid sending password hash
	newUser.Password = ""

	// creating permission group for user
	permissionGroup, err := c.permissionGroupRepo.Create(ctx, &models.PermissionGroup{
		Name: req.Username,
	})
	if err != nil {
		return nil, err
	}

	userPermission, err := c.userPermissionRepo.Create(ctx, &models.UserPermission{
		PermissionGroupID: permissionGroup.ID,
		UserID:            newUser.ID,
		Level:             models.AccessLevelOwner,
	})
	if err != nil {
		return nil, err
	}

	wallet, err := c.walletRepo.Create(ctx, &models.Wallet{
		Name:              req.Username + "_salary",
		Description:       "Salary wallet",
		PermissionGroupID: permissionGroup.ID,
		Currency:          models.Currency(req.SalaryCurrency),
		BankID:            bank.ID,
		IsSalary:          true,
	})
	if err != nil {
		return nil, err
	}

	var operationGroup *models.OperationGroup

	var operation *models.Operation

	if req.Salary != 0 {
		operationGroup, err = c.operationGroupRepo.Create(ctx, &models.OperationGroup{
			Name:     req.Username + "_salary",
			WalletID: wallet.ID,
		})
		if err != nil {
			return nil, err
		}

		operation, err = c.operationsRepo.Create(ctx, &models.Operation{
			OperationGroupID: operationGroup.ID,
			Time:             req.SalaryDate,
			IsMonthly:        true,
			Amount:           req.Salary,
			InitiatorID:      newUser.ID,
		})
		if err != nil {
			return nil, err
		}
	}

	return &responses.RegisterResponse{
		User:            newUser,
		PermissionGroup: permissionGroup,
		UserPermission:  userPermission,
		Wallet:          wallet,
		OperationGroup:  operationGroup,
		Operation:       operation,
	}, nil
}
