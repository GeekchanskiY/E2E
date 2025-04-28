package banks

import (
	"context"

	"finworker/internal/models"
)

func (r *repository) Create(ctx context.Context, bank *models.Bank) (*models.Bank, error) {
	q := `INSERT INTO banks(name) VALUES (:name) returning id`

	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &bank.ID, &bank)

	return bank, err
}
