package work

import (
	"context"
	"embed"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/models"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/works"
	"finworker/internal/templates"
)

type Controller interface {
	CreateWork(ctx context.Context) (*template.Template, map[string]any, error)
	CreateWorkForm(ctx context.Context, workTime *models.UserWork) (*template.Template, map[string]any, error)
	WorkTime(ctx context.Context) (*template.Template, map[string]any, error)
	WorkTimeForm(ctx context.Context, workId, userId int64) (*template.Template, map[string]any, error)
}

type controller struct {
	logger *zap.Logger

	userRepo *users.Repository
	workRepo *works.Repository

	secret string

	fs embed.FS
}

func New(
	logger *zap.Logger,

	userRepo *users.Repository,
	workRepo *works.Repository,

	secret string,
) Controller {
	return &controller{
		logger: logger,

		userRepo: userRepo,
		workRepo: workRepo,

		secret: secret,

		fs: templates.Fs,
	}
}
