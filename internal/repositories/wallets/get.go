package wallets

import (
	"context"

	"finworker/internal/models"
)

func (repo *repository) Get(ctx context.Context, id int64) (models.Wallet, error) {
	var wallet models.Wallet
	err := repo.db.GetContext(ctx, &wallet, `SELECT * FROM wallets WHERE id = $1`, id)

	return wallet, err
}
