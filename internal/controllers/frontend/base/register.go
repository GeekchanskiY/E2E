package base

import (
	"context"
	"database/sql"
	"errors"
	"html/template"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"
	utils2 "finworker/internal/utils"

	"go.uber.org/zap"
)

func (c *controller) Register(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))

	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *controller) RegisterForm(ctx context.Context, username, name, password, repeatPassword, gender, birthday, bank, salary, currency, payday string) (*template.Template, map[string]any, string, string, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))

	var (
		err error
	)

	if username == "" {
		return c.registerFormError(ctx, errors.New("username is required"))
	}

	if password == "" {
		return c.registerFormError(ctx, errors.New("password is required"))
	}

	if password != repeatPassword {
		return c.registerFormError(ctx, errors.New("password does not match"))
	}

	if gender != "male" && gender != "female" {
		return c.registerFormError(ctx, errors.New("gender is invalid"))
	}

	birthdayDate, err := time.Parse(time.DateOnly, birthday)
	if err != nil {
		return c.registerFormError(ctx, errors.New("birthday is invalid"))
	}

	if bank == "" {
		return c.registerFormError(ctx, errors.New("bank is required"))
	}

	if salary == "" {
		return c.registerFormError(ctx, errors.New("salary is required"))
	}

	salaryInt, err := strconv.Atoi(salary)
	if err != nil {
		return c.registerFormError(ctx, errors.New("salary is invalid"))
	}

	if currency == "" {
		return c.registerFormError(ctx, errors.New("currency is required"))
	}

	if payday == "" {
		return c.registerFormError(ctx, errors.New("payday is required"))
	}

	paydayInt, err := strconv.Atoi(payday)
	if err != nil {
		return c.registerFormError(ctx, errors.New("payday is invalid"))
	}

	dbBank, err := c.banksRepo.GetByName(ctx, bank)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.registerFormError(ctx, errors.New("bank does not exist"))
		}

		return c.registerFormError(ctx, err)
	}

	hashedPassword, err := utils2.HashPassword(password)
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	newUser := &models.User{
		Username: username,
		Password: hashedPassword,
		Name:     name,
		Gender:   gender,
		Birthday: birthdayDate,
	}

	newUser, err = c.userRepo.Create(ctx, newUser)
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	// sets password to empty to avoid sending password hash
	newUser.Password = ""

	// creating a permission group for user
	permissionGroup, err := c.permissionGroupsRepo.Create(ctx, &models.PermissionGroup{
		Name: username + "_group",
	})
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	_, err = c.userPermissionsRepo.Create(ctx, &models.UserPermission{
		PermissionGroupID: permissionGroup.ID,
		UserID:            newUser.ID,
		Level:             models.AccessLevelOwner,
	})
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	wallet, err := c.walletsRepo.Create(ctx, &models.Wallet{
		Name:              username + "_salary",
		Description:       "Salary wallet",
		PermissionGroupID: permissionGroup.ID,
		Currency:          models.Currency(currency),
		BankID:            dbBank.ID,
		IsSalary:          true,
	})
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	var operationGroup *models.OperationGroup

	if salaryInt != 0 {
		operationGroup, err = c.operationGroupsRepo.Create(ctx, &models.OperationGroup{
			Name:     username + "_salary",
			WalletID: wallet.ID,
		})
		if err != nil {
			return c.registerFormError(ctx, err)
		}

		_, err = c.operationsRepo.Create(ctx, &models.Operation{
			OperationGroupID: operationGroup.ID,
			Time:             time.Date(time.Now().Year(), time.Now().Month(), paydayInt, 0, 0, 0, 0, time.Local),
			IsMonthly:        true,
			Amount:           float64(salaryInt),
			InitiatorID:      newUser.ID,
		})
		if err != nil {
			return c.registerFormError(ctx, err)
		}
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": newUser,
		"id":   newUser.ID,
		"time": time.Now(),
	}).SignedString([]byte(c.secret))
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	salt, err := utils2.GenerateSaltFromPassword(password)
	if err != nil {
		return c.registerFormError(ctx, err)
	}

	return nil, nil, token, salt, nil
}

func (c *controller) registerFormError(ctx context.Context, userErr error) (*template.Template, map[string]any, string, string, error) {
	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, "", "", err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)
	data["error"] = userErr.Error()

	return html, data, "", "", err
}
