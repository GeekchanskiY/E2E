package wallets

import (
	"context"

	"finworker/internal/models"
)

func (repo *repository) Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	q := `
		INSERT INTO 
    	wallets (name, description, permission_group_id, currency, is_salary, bank_id) 
		VALUES (:name, :description, :permission_group_id, :currency, :is_salary, :bank_id) 
		returning id, created_at`

	namedStmt, err := repo.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.QueryRowxContext(ctx, wallet).Scan(&wallet.ID, &wallet.CreatedAt)

	return wallet, err
}
