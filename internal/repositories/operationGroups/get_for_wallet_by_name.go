package operationGroups

import (
	"context"
	"database/sql"
	"errors"

	"finworker/internal/models"
)

func (r *Repository) GetOrCreateForWalletByName(ctx context.Context, walletID int64, name string) (*models.OperationGroup, error) {
	var (
		operationGroup = new(models.OperationGroup)

		err error
	)

	q := `SELECT id, name, wallet_id FROM operation_groups where wallet_id = $1 and  name = $2`
	err = r.db.QueryRowxContext(ctx, q, walletID, name).StructScan(operationGroup)
	if err == nil {

		return operationGroup, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		if operationGroup, err = r.Create(ctx, &models.OperationGroup{
			Name:     name,
			WalletID: walletID,
		}); err != nil {
			return nil, err
		}

		return operationGroup, nil
	}

	return nil, err
}
