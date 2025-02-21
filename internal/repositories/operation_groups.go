package repositories

import (
	"github.com/jmoiron/sqlx"
)

type OperationGroupRepository struct {
	db *sqlx.DB
}

func NewOperationGroupRepository(db *sqlx.DB) *OperationGroupRepository {
	return &OperationGroupRepository{db: db}
}
