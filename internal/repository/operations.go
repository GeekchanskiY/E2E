package repository

import (
	"github.com/jmoiron/sqlx"
)

type OperationRepository struct {
	db *sqlx.DB
}

func NewOperationRepository(db *sqlx.DB) *OperationRepository {
	return &OperationRepository{db: db}
}
