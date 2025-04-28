package operationGroups

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, operationGroup *models.OperationGroup) (*models.OperationGroup, error)
	GetByWallet(ctx context.Context, walletID int64) (operationGroups []*models.OperationGroup, err error)
	GetOrCreateForWalletByName(ctx context.Context, walletID int64, name string) (*models.OperationGroup, error)
}

type repository struct {
	log *zap.Logger
	db  *sqlx.DB
}

func New(db *sqlx.DB, log *zap.Logger) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}
