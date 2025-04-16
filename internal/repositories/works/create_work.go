package works

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) CreateUserWork(ctx context.Context, work *models.UserWork) (*models.UserWork, error) {
	if err := r.db.GetContext(ctx, &work.ID, `INSERT INTO user_work(name, hourly_rate, worker) VALUES ($1, $2, $3) returning id`, work.Name, work.HourlyRate, work.Worker); err != nil {
		return nil, err
	}

	return work, nil
}
