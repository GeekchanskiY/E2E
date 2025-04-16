package currencyStates

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository struct {
	log *zap.Logger
	db  *sqlx.DB
}

func New(db *sqlx.DB, log *zap.Logger) *Repository {
	return &Repository{
		log: log,
		db:  db,
	}
}
