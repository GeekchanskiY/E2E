package frontend

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"

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

func (c *Controller) RegisterForm(ctx context.Context, username, password, repeatPassword, gender, birthday, bank, salary, currency, payday string) (*template.Template, map[string]any, string, string, error) {
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

	data["error"] = "form not ready"
	return html, data, "", "", errors.New("form not ready")
}
