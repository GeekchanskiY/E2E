package distributors

import (
	"context"

	"go.uber.org/zap"

	"finworker/internal/models"
)

func (repo *Repository) GetForWallet(ctx context.Context, walletID int64) ([]*models.DistributorExtended, error) {
	q := `SELECT 
    distributors.id, distributors.name, distributors.source_wallet_id, source.name, distributors.target_wallet_id, target.name, distributors.percent
	FROM distributors
	join wallets as source on distributors.source_wallet_id = source.id
	join wallets as target on distributors.target_wallet_id = target.id
    WHERE source_wallet_id = $1`

	rows, err := repo.db.QueryxContext(ctx, q, walletID)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			repo.log.Error("rows close error", zap.Error(err))
		}
	}()

	var distributors = make([]*models.DistributorExtended, 0)

	for rows.Next() {
		distributor := new(models.DistributorExtended)

		err = rows.Scan(
			&distributor.ID,
			&distributor.Name,
			&distributor.SourceWalletID,
			&distributor.SourceWalletName,
			&distributor.TargetWalletID,
			&distributor.TargetWalletName,
			&distributor.Percent,
		)

		if err != nil {
			return nil, err
		}

		distributors = append(distributors, distributor)
	}

	return distributors, err
}
