package works

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	CreateUserWork(ctx context.Context, work *models.UserWork) (*models.UserWork, error)
	EndWorkTime(ctx context.Context, workID int64) error
	StartWorkTime(ctx context.Context, workID int64) (*models.WorkTime, error)
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
