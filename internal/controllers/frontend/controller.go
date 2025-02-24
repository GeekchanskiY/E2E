package frontend

import (
	"embed"

	"finworker/internal/templates"
	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger

	fs embed.FS
}

func New(logger *zap.Logger) *Controller {

	return &Controller{
		logger: logger,
		fs:     templates.Fs,
	}
}
