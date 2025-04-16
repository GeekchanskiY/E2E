package operation_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) GetByWallet(ctx context.Context, walletId int64) (operationGroups []*models.OperationGroup, err error) {

	q := `SELECT id, name, wallet_id FROM operation_groups where wallet_id = $1`

	rows, err := r.db.QueryxContext(ctx, q, walletId)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
		}
	}()

	for rows.Next() {
		var operationGroup models.OperationGroup

		err = rows.Scan(&operationGroup.Id, &operationGroup.Name, &operationGroup.WalletId)
		if err != nil {
			return nil, err
		}

		operationGroups = append(operationGroups, &operationGroup)
	}

	return operationGroups, err
}
