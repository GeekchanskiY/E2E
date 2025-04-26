package users

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Get(ctx context.Context, id int) (models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
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
