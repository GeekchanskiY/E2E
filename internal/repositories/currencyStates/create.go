package currencyStates

import (
	"context"

	"finworker/internal/models"
)

func (r *Repository) Create(ctx context.Context, currencyState *models.CurrencyState) (*models.CurrencyState, error) {
	q := `INSERT INTO currency_states(bank_id, source_name, currency_name, sell_usd, buy_usd, time) 
		VALUES (:bank_id, :source_name, :currency_name, :sell_usd, :buy_usd, :time) returning id`

	namedStmt, err := r.db.PrepareNamed(q)
	if err != nil {
		return nil, err
	}

	err = namedStmt.GetContext(ctx, &currencyState.ID, currencyState)

	return currencyState, err
}
