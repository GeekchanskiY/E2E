package wallets

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error)
	Get(ctx context.Context, id int64) (models.Wallet, error)
	GetByUsername(ctx context.Context, username string) ([]models.WalletExtended, error)
}

type repository struct {
	db *sqlx.DB

	log *zap.Logger
}

func New(db *sqlx.DB, log *zap.Logger) Repository {
	return &repository{
		db:  db,
		log: log,
	}
}
