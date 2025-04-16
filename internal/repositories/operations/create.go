package operations

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) Create(ctx context.Context, operation *models.Operation) (*models.Operation, error) {
	q := `INSERT INTO operations(operation_group_id, amount, time, is_monthly, is_confirmed, initiator_id)
		VALUES (:operation_group_id, :amount, :time, :is_monthly, :is_confirmed, :initiator_id) returning id`
	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &operation.ID, operation)

	return operation, err
}
