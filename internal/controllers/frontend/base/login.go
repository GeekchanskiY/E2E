package base

import (
	"database/sql"
	"errors"
	"html/template"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
	"finworker/internal/utils"
)

func (c *controller) Login(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.login.controller", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.LoginTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *controller) LoginForm(ctx context.Context, username, password string) (*template.Template, map[string]any, string, string, error) {
	c.logger.Debug("frontend.login.controller.form", zap.String("event", "got request"))

	html, err := templateUtils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.LoginTemplate)
	if err != nil {
		return nil, nil, "", "", err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	user, err := c.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			data["error"] = "user not found"
			return html, data, "", "", errors.New("user not found")
		}

		data["error"] = err.Error()
		return html, data, "", "", err
	}

	isPasswordCorrect := utils.VerifyPassword(password, user.Password)
	if !isPasswordCorrect {
		data["error"] = "invalid password"
		return html, data, "", "", errors.New("invalid password")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"id":   user.ID,
		"time": time.Now(),
	}).SignedString([]byte(c.secret))
	if err != nil {
		data["error"] = err.Error()
		return html, data, "", "", err
	}
	salt, err := utils.GenerateSaltFromPassword(password)
	if err != nil {
		data["error"] = err.Error()
		return html, data, "", "", err
	}

	return html, data, token, salt, nil
}
