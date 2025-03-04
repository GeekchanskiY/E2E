package frontend

import (
	"context"
	"database/sql"
	"errors"
	"html/template"
	"time"

	"finworker/internal/templates"
	"finworker/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

func (c *Controller) Login(_ context.Context) (*template.Template, error) {
	c.logger.Debug("frontend.login.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.LoginTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}

func (c *Controller) LoginForm(ctx context.Context, username, password string) (*template.Template, string, error) {
	c.logger.Debug("frontend.login.controller.form", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.LoginTemplate)
	if err != nil {
		return nil, "", err
	}

	user, err := c.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return html, "", errors.New("user not found")
		}
		return html, "", err
	}

	isPasswordCorrect := utils.VerifyPassword(password, user.Password)
	if !isPasswordCorrect {
		return html, "", errors.New("wrong password")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"id":   user.Id,
		"time": time.Now(),
	}).SignedString([]byte(c.secret))
	if err != nil {
		return html, "", err
	}

	return html, token, nil
}
