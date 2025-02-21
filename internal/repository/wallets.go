package repository

import (
	"context"
	"fmt"

	"finworker/internal/models"
	"github.com/jmoiron/sqlx"
)

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (w *WalletRepository) CreateWallet(ctx *context.Context, wallet *models.Wallet) (*models.Wallet, error) {
	q := `INSERT INTO wallets (name, description, permission_group_id, created_at, currency, is_salary) VALUES (:name, :description, :permission_group_id, :created_at, :currency, :is_salary)`
	fmt.Println(q)
	return nil, nil
}
