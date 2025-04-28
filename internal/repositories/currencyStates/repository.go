package currencyStates

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, currencyState *models.CurrencyState) (*models.CurrencyState, error)
	GetLastUpdate(ctx context.Context) (time.Time, error)
	GetBankCurrencyState(ctx context.Context, currencyCode models.Currency, bankID int64) (*models.CurrencyState, error)
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
