package distributors

import (
	"context"

	"finworker/internal/models"
)

func (repo *Repository) GetForWallet(ctx context.Context, walletId int) ([]models.DistributorExtended, error) {
	q := `SELECT 
    distributors.id, distributors.name, distributors.source_wallet_id, source.name, distributors.target_wallet_id, target.name, distributors.percent
	FROM distributors
	join wallets as source on distributors.source_wallet_id = source.id
	join wallets as target on distributors.target_wallet_id = target.id
    WHERE source_wallet_id = $1`

	rows, err := repo.db.QueryxContext(ctx, q, walletId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
		}
	}()

	var distributors = make([]models.DistributorExtended, 0)
	for rows.Next() {
		var distributor models.DistributorExtended

		err = rows.Scan(
			&distributor.Id,
			&distributor.Name,
			&distributor.SourceWalletId,
			&distributor.SourceWalletName,
			&distributor.TargetWalletId,
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
