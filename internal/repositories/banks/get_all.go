package banks

import (
	"go.uber.org/zap"

	"finworker/internal/models"
)

func (r *Repository) GetAll() (banks []*models.Bank, err error) {
	q := `SELECT id, name FROM banks`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			r.log.Error("rows close error", zap.Error(err))
		}
	}()

	for rows.Next() {
		var bank models.Bank

		err = rows.Scan(&bank.ID, &bank.Name)
		if err != nil {
			return nil, err
		}

		banks = append(banks, &bank)
	}

	return banks, err
}
