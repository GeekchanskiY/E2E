package permission_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *repository) GetByName(ctx context.Context, name string) (*models.PermissionGroup, error) {
	q := `SELECT id, name, created_at, updated_at FROM permission_groups WHERE name=$1`

	group := models.PermissionGroup{}
	if err := r.db.QueryRowxContext(ctx, q, name).Scan(&group.ID, &group.Name, &group.CreatedAt, &group.UpdatedAt); err != nil {
		return nil, err
	}

	return &group, nil
}
