package media

import (
	"context"
	"mime/multipart"

	"go.uber.org/zap"
)

const mediaPath = "media"

type Controller interface {
	Upload(ctx context.Context, fileHeader *multipart.FileHeader, file []byte) error
}

type controller struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) Controller {
	return &controller{
		logger: logger,
	}
}
