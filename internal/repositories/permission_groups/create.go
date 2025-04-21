package permission_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) Create(ctx context.Context, group *models.PermissionGroup) (*models.PermissionGroup, error) {
	q := `INSERT INTO permission_groups(name) VALUES (:name) returning id, created_at, updated_at`

	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.QueryRowxContext(ctx, group).Scan(&group.ID, &group.CreatedAt, &group.UpdatedAt)

	return group, err
}
