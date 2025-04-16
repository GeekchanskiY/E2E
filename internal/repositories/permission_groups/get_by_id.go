package permission_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) GetByID(ctx context.Context, id int64) (*models.PermissionGroup, error) {
	q := `SELECT id, name, created_at, updated_at FROM permission_groups WHERE id=$1`

	group := models.PermissionGroup{}
	if err := r.db.QueryRowxContext(ctx, q, id).Scan(&group.ID, &group.Name, &group.CreatedAt, &group.UpdatedAt); err != nil {
		return nil, err
	}

	return &group, nil
}
