package banks

import (
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
		}
	}()

	for rows.Next() {
		var bank models.Bank

		err = rows.Scan(&bank.Id, &bank.Name)
		if err != nil {
			return nil, err
		}

		banks = append(banks, &bank)
	}

	return banks, err
}
