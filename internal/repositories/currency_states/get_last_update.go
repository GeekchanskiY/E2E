package currency_states

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (r *Repository) GetLastUpdate(ctx context.Context) (time.Time, error) {
	q := `SELECT time FROM currency_states ORDER BY time DESC LIMIT 1`

	var lastTime time.Time
	err := r.db.QueryRowxContext(ctx, q).Scan(&lastTime)

	if errors.Is(err, sql.ErrNoRows) {
		return time.Time{}, nil
	}

	return lastTime, err
}
