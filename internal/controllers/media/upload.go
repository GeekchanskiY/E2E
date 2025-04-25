package media

import (
	"context"
	"mime/multipart"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func (c *controller) Upload(ctx context.Context, fileHeader *multipart.FileHeader, file []byte) error {
	var (
		dst *os.File
		err error
	)
	if _, err = os.Stat(mediaPath); os.IsNotExist(err) {
		if err = os.Mkdir(mediaPath, 0755); err != nil {
			return err
		}
	}

	if dst, err = os.Create(filepath.Join(mediaPath, fileHeader.Filename)); err != nil {
		return err
	}

	defer func() {
		if err = dst.Close(); err != nil {
			c.logger.Fatal("media.upload.controller failed to close dst", zap.Error(err))
		}
	}()

	if _, err = dst.Write(file); err != nil {
		return err
	}

	return nil
}
