package repositories

import (
	"context"

	"finworker/internal/models"
	"github.com/jmoiron/sqlx"
)

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(ctx context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	q := `
		INSERT INTO 
    	wallets (name, description, permission_group_id, currency, is_salary, bank_id) 
		VALUES (:name, :description, :permission_group_id, :currency, :is_salary, :bank_id) 
		returning id, created_at`

	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.QueryRowxContext(ctx, wallet).Scan(&wallet.Id, &wallet.CreatedAt)

	return wallet, err
}
