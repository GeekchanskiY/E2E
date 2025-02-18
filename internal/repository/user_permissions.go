package repository

import (
	"github.com/jmoiron/sqlx"
)

type UserPermissionRepository struct {
	db *sqlx.DB
}

func NewUserPermissionRepository(db *sqlx.DB) *UserPermissionRepository {
	return &UserPermissionRepository{db: db}
}
