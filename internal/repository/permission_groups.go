package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"finworker/internal/models"
)

type PermissionGroupRepository struct {
	db *sqlx.DB
}

func NewPermissionGroupRepository(db *sqlx.DB) *PermissionGroupRepository {
	return &PermissionGroupRepository{db: db}
}

func (r *PermissionGroupRepository) Create(ctx context.Context, group *models.PermissionGroup) (*models.PermissionGroup, error) {
	q := `INSERT INTO permission_groups(name, level) VALUES (:name, :level) returning id`
	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &group.Id, &group)

	return group, err
}
