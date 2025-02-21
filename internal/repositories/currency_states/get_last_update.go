package currency_states

import (
	"context"
	"time"
)

func (r *Repository) GetLastUpdate(ctx context.Context) (*time.Time, error) {
	q := `SELECT time FROM currency_states ORDER BY time DESC LIMIT 1`

	var lastTime *time.Time
	err := r.db.GetContext(ctx, lastTime, q)

	return lastTime, err
}
