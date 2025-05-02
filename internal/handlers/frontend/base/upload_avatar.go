package base

import (
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"finworker/internal/config"
)

func (h *handler) UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var (
		file   multipart.File
		header *multipart.FileHeader

		fileBytes []byte
		fileType  string

		err error
	)

	// Limit uploads to 10MB
	if err = r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "file is too big!", http.StatusBadRequest)

		return
	}

	if file, header, err = r.FormFile("file"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	h.logger.Info(
		"media.upload",
		zap.String("filename", header.Filename),
		zap.String("mimetype", header.Header.Get("Content-Type")),
		zap.Int64("size", header.Size),
	)

	if fileBytes, err = io.ReadAll(file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	fileType = http.DetectContentType(fileBytes)
	if !strings.HasPrefix(fileType, "image/") {
		http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)

		return
	}

	if err = h.controller.UploadAvatar(r.Context(), r.Context().Value(config.UserIDContextKey).(int64), header, fileBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/me", http.StatusSeeOther)

	return
}
