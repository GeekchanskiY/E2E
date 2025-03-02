package frontend

import (
	"context"
	"database/sql"
	"errors"
	"html/template"

	"finworker/internal/utils"
)

func (c *Controller) Login(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.login")

	html, err := template.ParseFS(c.fs, "base.gohtml", "login.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}

func (c *Controller) LoginForm(ctx context.Context, username, password string) (*template.Template, error) {
	c.logger.Info("frontend.login.form")

	html, err := template.ParseFS(c.fs, "base.gohtml", "login.gohtml")
	if err != nil {
		return nil, err
	}

	user, err := c.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return html, errors.New("user not found")
		}
		return html, err
	}

	isPasswordCorrect := utils.VerifyPassword(password, user.Password)
	if !isPasswordCorrect {
		return html, errors.New("wrong password")
	}

	return html, nil
}
