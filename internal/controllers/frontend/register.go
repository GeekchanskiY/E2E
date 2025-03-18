package frontend

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (c *Controller) Register(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *Controller) RegisterForm(ctx context.Context, username, name, password, repeatPassword, gender, birthday, bank, salary, currency, payday string) (*template.Template, map[string]any, string, string, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))

	var (
		err error
	)

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, "", "", err
	}
	data := utils.BuildDefaultDataMapFromContext(ctx)

	if username == "" {
		err = errors.New("username is required")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if password == "" {
		err = errors.New("password is required")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if password != repeatPassword {
		err = errors.New("password does not match")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if gender != "male" && gender != "female" {
		err = errors.New("gender is invalid")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	birthdayDate, err := time.Parse(time.DateOnly, birthday)
	if err != nil {
		err = errors.New("birthday is invalid")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if bank == "" {
		err = errors.New("bank is required")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if salary == "" {
		err = errors.New("salary is required")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	salaryInt, err := strconv.Atoi(salary)
	if err != nil {
		err = errors.New("salary is invalid")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if currency == "" {
		err = errors.New("currency is required")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	if payday == "" {
		err = errors.New("payday is required")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	paydayInt, err := strconv.Atoi(payday)
	if err != nil {
		err = errors.New("payday is invalid")
		data["error"] = err.Error()

		return html, data, "", "", err
	}

	fmt.Println("birthdayDate:", birthdayDate)
	fmt.Println("salaryInt:", salaryInt)
	fmt.Println("currency:", paydayInt)

	dbBank, err := c.banksRepo.GetByName(bank)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = errors.New("bank does not exist")
			data["error"] = err.Error()
			return html, data, "", "", err
		}

		data["error"] = err.Error()
		return html, data, "", "", err
	}

	hashedPassword, err := utils2.HashPassword(password)
	if err != nil {
		err = errors.New("failed to hash password")
		data["error"] = err.Error()

		return html, data, "", "", err
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
		data["error"] = err.Error()
		return html, data, "", "", err
	}

	// sets password to empty to avoid sending password hash
	newUser.Password = ""

	// creating permission group for user
	permissionGroup, err := c.permissionGroupsRepo.Create(ctx, &models.PermissionGroup{
		Name: username + "_group",
	})
	if err != nil {
		data["error"] = err.Error()
		return html, data, "", "", err
	}

	_, err = c.userPermissionsRepo.Create(ctx, &models.UserPermission{
		PermissionGroupId: permissionGroup.Id,
		UserId:            newUser.Id,
		Level:             models.AccessLevelOwner,
	})
	if err != nil {
		return html, data, "", "", err
	}

	wallet, err := c.walletsRepo.Create(ctx, &models.Wallet{
		Name:              username + "_salary",
		Description:       "Salary wallet",
		PermissionGroupId: permissionGroup.Id,
		Currency:          models.CurrencyBYN,
		BankId:            dbBank.Id,
		IsSalary:          true,
	})
	if err != nil {
		return html, data, "", "", err
	}

	var operationGroup *models.OperationGroup

	if salaryInt != 0 {
		operationGroup, err = c.operationGroupsRepo.Create(ctx, &models.OperationGroup{
			Name:     username + "_salary",
			WalletId: wallet.Id,
		})
		if err != nil {
			data["error"] = err.Error()
			return html, data, "", "", err
		}

		_, err = c.operationsRepo.Create(ctx, &models.Operation{
			OperationGroupId: operationGroup.Id,
			IsConsumption:    false,
			Time:             time.Date(time.Now().Year(), time.Now().Month(), paydayInt, 0, 0, 0, 0, time.Local),
			IsMonthly:        true,
			Amount:           float64(salaryInt),
			InitiatorId:      newUser.Id,
		})
		if err != nil {
			data["error"] = err.Error()
			return html, data, "", "", err
		}

	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": newUser,
		"id":   newUser.Id,
		"time": time.Now(),
	}).SignedString([]byte(c.secret))
	if err != nil {
		data["error"] = err.Error()
		return html, data, "", "", err
	}
	salt, err := utils2.GenerateSaltFromPassword(password)
	if err != nil {
		data["error"] = err.Error()
		return html, data, "", "", err
	}

	return html, data, token, salt, nil
}
