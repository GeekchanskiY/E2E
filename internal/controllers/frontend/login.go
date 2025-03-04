package frontend

import (
	"context"
	"database/sql"
	"errors"
	"html/template"
	"time"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
	"finworker/internal/utils"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

func (c *Controller) Login(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.login.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.LoginTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *Controller) LoginForm(ctx context.Context, username, password string) (*template.Template, map[string]any, string, error) {
	c.logger.Debug("frontend.login.controller.form", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.LoginTemplate)
	if err != nil {
		return nil, nil, "", err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	user, err := c.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return html, data, "", errors.New("user not found")
		}
		return html, data, "", err
	}

	isPasswordCorrect := utils.VerifyPassword(password, user.Password)
	if !isPasswordCorrect {
		return html, data, "", errors.New("wrong password")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"id":   user.Id,
		"time": time.Now(),
	}).SignedString([]byte(c.secret))
	if err != nil {
		return html, data, "", err
	}

	return html, data, token, nil
}
