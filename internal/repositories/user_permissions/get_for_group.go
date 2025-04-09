package user_permissions

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) GetForGroup(ctx context.Context, groupId int64) ([]*models.UserPermissionExtended, error) {
	q := `
	SELECT user_permission.id, permission_group_id, user_id, level, created_at, users.username
	FROM user_permission
	JOIN users ON user_permission.user_id = users.id
	WHERE permission_group_id = $1
	`

	rows, err := r.db.QueryxContext(ctx, q, groupId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			panic(err)
		}
	}()

	userPermissions := make([]*models.UserPermissionExtended, 0)
	for rows.Next() {
		userPermission := new(models.UserPermissionExtended)
		if err := rows.StructScan(&userPermission); err != nil {
			return nil, err
		}

		userPermissions = append(userPermissions, userPermission)
	}

	return userPermissions, nil
}
