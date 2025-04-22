package permission_groups

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"finworker/internal/models"
)

func (r *Repository) GetUserEditGroups(ctx context.Context, userID int64) (permissionGroups []*models.PermissionGroup, err error) {
	var (
		q    string
		rows *sqlx.Rows
	)

	q = `
	select
    	permission_groups.id, permission_groups.name, permission_groups.created_at, permission_groups.updated_at
	from permission_groups 
    	join user_permission on permission_groups.id = user_permission.permission_group_id
    where
        user_permission.user_id = $1
    and 
        (user_permission.level = 'owner' or user_permission.level = 'full')`

	if rows, err = r.db.QueryxContext(ctx, q, userID); err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			r.log.Error("failed to close rows", zap.Error(err))
		}
	}()

	for rows.Next() {
		var permissionGroup models.PermissionGroup

		if err = rows.Scan(&permissionGroup.ID, &permissionGroup.Name, &permissionGroup.CreatedAt, &permissionGroup.UpdatedAt); err != nil {
			return nil, err
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}

	return permissionGroups, err
}

func (r *Repository) GetUserGroups(ctx context.Context, userID int64) (permissionGroups []*models.PermissionGroupWithRole, err error) {
	var (
		rows *sqlx.Rows
	)

	q := `
	select
    	permission_groups.id, permission_groups.name, permission_groups.created_at, permission_groups.updated_at, user_permission.level,
    	(select 
    	     count(*) 
    	from 
    	    user_permission 
    	where 
    	    permission_group_id = permission_groups.id
    	) as users_count
	from permission_groups
		join user_permission on user_permission.permission_group_id = permission_groups.id
	where
		user_permission.user_id = $1;
    `

	if rows, err = r.db.QueryxContext(ctx, q, userID); err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			r.log.Error("failed to close rows", zap.Error(err))
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
