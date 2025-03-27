package finance

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/finance"
)

type Handler interface {
	Wallet(w http.ResponseWriter, r *http.Request)
	Finance(w http.ResponseWriter, r *http.Request)
	CreateWallet(w http.ResponseWriter, r *http.Request)
	CreateOperationGroup(w http.ResponseWriter, r *http.Request)
	CreateOperation(w http.ResponseWriter, r *http.Request)
	CreateDistributor(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger     *zap.Logger
	controller finance.Controller
}

func New(logger *zap.Logger, controller finance.Controller) Handler {
	return &handler{
		logger:     logger,
		controller: controller,
	}
}
