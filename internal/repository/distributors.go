package repository

import (
	"github.com/jmoiron/sqlx"
)

type DistributorsRepository struct {
	db *sqlx.DB
}

func NewDistributorsRepository(db *sqlx.DB) *DistributorsRepository {
	return &DistributorsRepository{db: db}
}
