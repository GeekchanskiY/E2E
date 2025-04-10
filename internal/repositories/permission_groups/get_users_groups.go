package permission_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) GetUserEditGroups(ctx context.Context, userId int64) (permissionGroups []*models.PermissionGroup, err error) {
	q := `SELECT
    	permission_groups.id, permission_groups.name, permission_groups.created_at, permission_groups.updated_at
	FROM permission_groups 
    	JOIN user_permission ON permission_groups.id = user_permission.user_id
    WHERE
        user_permission.user_id = $1
    AND 
        (user_permission.level = 'owner' or user_permission.level = 'full')`

	rows, err := r.db.QueryContext(ctx, q, userId)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
		}
	}()

	for rows.Next() {
		var permissionGroup models.PermissionGroup
		err = rows.Scan(&permissionGroup.Id, &permissionGroup.Name, &permissionGroup.CreatedAt, &permissionGroup.UpdatedAt)
		if err != nil {
			return nil, err
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}

	return permissionGroups, err
}

func (r *Repository) GetUserGroups(ctx context.Context, userId int64) (permissionGroups []*models.PermissionGroupWithRole, err error) {
	q := `
	SELECT
    permission_groups.id, permission_groups.name, permission_groups.created_at, permission_groups.updated_at, user_permission.level,
    (select count(*) from user_permission where permission_group_id = permission_groups.id) as users_count
	FROM permission_groups
			 JOIN user_permission ON user_permission.permission_group_id = permission_groups.id
	WHERE
		user_permission.user_id = $1;
    `

	rows, err := r.db.QueryxContext(ctx, q, userId)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
		}
	}()

	for rows.Next() {
		permissionsGroup := new(models.PermissionGroupWithRole)
		err = rows.StructScan(permissionsGroup)
		if err != nil {
			return nil, err
		}

		permissionGroups = append(permissionGroups, permissionsGroup)
	}

	return permissionGroups, err
}
