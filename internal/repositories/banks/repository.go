package banks

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, bank *models.Bank) (*models.Bank, error)
	GetAll() (banks []*models.Bank, err error)
	GetByID(id int64) (*models.Bank, error)
	GetByName(ctx context.Context, name string) (*models.Bank, error)
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
