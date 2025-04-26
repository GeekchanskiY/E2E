package registry

import (
	"context"

	"finworker/internal/models"
)

func (r *repository) Push(ctx context.Context, event *models.Event) (*models.Event, error) {
	q := `insert into registry(name, event, content) values ($1, $2, $3) returning id, time`

	return event, r.db.QueryRowxContext(ctx, q, event.Name, event.Event, event.Content).Scan(&event.ID, &event.Time)
}
