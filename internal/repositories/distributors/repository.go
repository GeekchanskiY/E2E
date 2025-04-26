package distributors

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, distributor *models.Distributor) (*models.Distributor, error)
	GetForWallet(ctx context.Context, walletID int64) ([]*models.DistributorExtended, error)
}

type repository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func New(db *sqlx.DB, log *zap.Logger) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}
