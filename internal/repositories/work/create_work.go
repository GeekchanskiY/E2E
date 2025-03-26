package work

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) CreateUserWork(ctx context.Context, wallet *models.UserWork) (*models.UserWork, error) {
	if err := r.db.GetContext(ctx, &wallet.Id, `INSERT INTO user_work(name, hourly_rate, worker) VALUES (:name, :hourly_rate, :worker) returning id`, wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}
