package repository

import (
	"github.com/jmoiron/sqlx"

	"finworker/internal/models"
)

type DistributorsRepository struct {
	db *sqlx.DB
}

func NewDistributorsRepository(db *sqlx.DB) *DistributorsRepository {
	return &DistributorsRepository{db: db}
}

func (r *DistributorsRepository) Insert(distributor *models.Distributor) error {
	return nil
}
