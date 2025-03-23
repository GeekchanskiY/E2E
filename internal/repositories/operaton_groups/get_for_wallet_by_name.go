package operaton_groups

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) GetForWalletByName(ctx context.Context, walletId int64, name string) (operationGroup *models.OperationGroup, err error) {

	q := `SELECT id, name, wallet_id FROM operation_groups where wallet_id = $1 and  name = $2`

	err = r.db.QueryRowxContext(ctx, q, walletId, name).StructScan(&operationGroup)
	if err != nil {
		return nil, err
	}

	return operationGroup, err
}
