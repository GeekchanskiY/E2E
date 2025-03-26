package work

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository struct {
	db *sqlx.DB

	log *zap.Logger
}

func New(db *sqlx.DB, log *zap.Logger) *Repository {
	return &Repository{
		db:  db,
		log: log,
	}
}
