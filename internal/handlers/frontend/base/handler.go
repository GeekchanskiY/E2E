package base

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/base"
)

type Handler interface {
	FAQ(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
	PageNotFound(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	UIKit(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger     *zap.Logger
	controller base.Controller
}

func New(logger *zap.Logger, controller base.Controller) Handler {
	return &handler{
		logger:     logger,
		controller: controller,
	}
}
