package repository

import (
	"github.com/jmoiron/sqlx"
)

type PermissionGroupRepository struct {
	db *sqlx.DB
}

func NewPermissionGroupRepository(db *sqlx.DB) *PermissionGroupRepository {
	return &PermissionGroupRepository{db: db}
}
