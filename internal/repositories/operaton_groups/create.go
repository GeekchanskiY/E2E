package operaton_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) Create(ctx context.Context, operationGroup *models.OperationGroup) (*models.OperationGroup, error) {
	q := `INSERT INTO operation_groups(name, wallet_id) 
		VALUES (:name, :wallet_id) returning id`
	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &operationGroup.Id, operationGroup)

	return operationGroup, err
}
