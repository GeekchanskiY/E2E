package banks

import (
	"finworker/internal/models"
)

func (r *Repository) GetById(id int64) (*models.Bank, error) {
	q := `SELECT id, name FROM banks WHERE id=$1`

	bank := new(models.Bank)
	err := r.db.QueryRow(q, id).Scan(&bank.Id, &bank.Name)

	return bank, err
}
