package works

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) StartWorkTime(ctx context.Context, workId int64) (*models.WorkTime, error) {
	var workTime *models.WorkTime
	if err := r.db.GetContext(ctx, workTime, `INSERT INTO work_time(work) VALUES ($1) returning id, work, start_time, end_time`, workId); err != nil {
		return nil, err
	}

	return workTime, nil
}
