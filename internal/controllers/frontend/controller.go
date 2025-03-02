package frontend

import (
	"embed"

	"finworker/internal/repositories/users"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger

	userRepo *users.Repository

	fs embed.FS
}

func New(logger *zap.Logger, userRepo *users.Repository) *Controller {

	return &Controller{
		logger: logger,

		userRepo: userRepo,

		fs: templates.Fs,
	}
}
