package banks

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) GetByName(ctx context.Context, name string) (*models.Bank, error) {
	q := `SELECT id, name FROM banks WHERE name=$1`

	bank := new(models.Bank)
	err := r.db.QueryRowxContext(ctx, q, name).Scan(&bank.Id, &bank.Name)

	return bank, err
}
