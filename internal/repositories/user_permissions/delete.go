package user_permissions

import (
	"context"
	"errors"
)

func (r *Repository) Delete(ctx context.Context, username string, permissionGroupId int64) error {
	tx, err := r.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}

	q := `delete from 
           user_permission 
       where user_id = (select id from users where username = $1) and permission_group_id = $2`

	_, err = tx.ExecContext(ctx, q, username, permissionGroupId)
	if err != nil {
		rollbackErr := tx.Rollback()
		return errors.Join(err, rollbackErr)
	}

	q = `UPDATE permission_groups SET updated_at = current_timestamp WHERE id = $2`
	_, err = tx.ExecContext(ctx, q, permissionGroupId)
	if err != nil {
		rollbackErr := tx.Rollback()
		return errors.Join(err, rollbackErr)
	}

	err = tx.Commit()

	return err
}
