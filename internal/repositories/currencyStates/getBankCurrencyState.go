package currencyStates

import (
	"context"

	"finworker/internal/models"
)

func (r *repository) GetBankCurrencyState(ctx context.Context, currencyCode models.Currency, bankID int64) (*models.CurrencyState, error) {
	currencyState := new(models.CurrencyState)

	q := `select * from currency_states WHERE currency_name = $1 and bank_id = $2 order by time desc limit 1`

	return currencyState, r.db.GetContext(ctx, currencyState, q, currencyCode, bankID)
}
