package operationGroups

import (
	"context"

	"go.uber.org/zap"

	"finworker/internal/models"
)

func (r *Repository) GetByWallet(ctx context.Context, walletID int64) (operationGroups []*models.OperationGroup, err error) {
	q := `SELECT id, name, wallet_id FROM operation_groups where wallet_id = $1`

	rows, err := r.db.QueryxContext(ctx, q, walletID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			r.log.Error("rows close error", zap.Error(err))
		}
	}()

	for rows.Next() {
		var operationGroup models.OperationGroup

		err = rows.Scan(&operationGroup.ID, &operationGroup.Name, &operationGroup.WalletID)
		if err != nil {
			return nil, err
		}

		operationGroups = append(operationGroups, &operationGroup)
	}

	return operationGroups, err
}
