package repositories

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
	q := `INSERT INTO permission_groups(name) VALUES (:name) returning id, created_at, updated_at`
	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.QueryRowxContext(ctx, group).Scan(&group.Id, &group.CreatedAt, &group.UpdatedAt)

	return group, err
}
