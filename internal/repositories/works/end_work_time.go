package works

import (
	"context"
	"errors"
)

func (r *Repository) EndWorkTime(ctx context.Context, workId int64) error {
	res, err := r.db.ExecContext(ctx, `UPDATE work_time SET end_time = current_timestamp where id = $1`, workId)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no active work time")
	}
	return nil
}
