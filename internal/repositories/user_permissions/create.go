package user_permissions

import (
	"context"
	"errors"
	"time"

	"finworker/internal/models"
)

func (r *Repository) Create(ctx context.Context, permission *models.UserPermission) (*models.UserPermission, error) {
	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	q := `INSERT INTO user_permission (permission_group_id, user_id, level) VALUES (:permission_group_id, :user_id, :level) RETURNING id, created_at`
	namedStmt, err := tx.PrepareNamed(q)
	if err != nil {
		rollbackErr := tx.Rollback()
		return nil, errors.Join(err, rollbackErr)
	}

	err = namedStmt.QueryRowxContext(ctx, permission).Scan(&permission.Id, &permission.CreatedAt)
	if err != nil {
		rollbackErr := tx.Rollback()
		return nil, errors.Join(err, rollbackErr)
	}

	q = `UPDATE permission_groups SET updated_at = $1 WHERE id = $2`
	_, err = tx.ExecContext(ctx, q, time.Now(), permission.Id)
	if err != nil {
		rollbackErr := tx.Rollback()
		return nil, errors.Join(err, rollbackErr)
	}

	err = tx.Commit()

	return permission, err
}
