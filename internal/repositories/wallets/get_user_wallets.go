package wallets

import (
	"context"

	"finworker/internal/models"

	"go.uber.org/zap"
)

func (repo *repository) GetByUsername(ctx context.Context, username string) ([]models.WalletExtended, error) {
	q := `SELECT 
    wallets.id, 
    wallets.name, 
    wallets.description,
    user_permission.level, 
    wallets.created_at,
    wallets.currency, 
    wallets.is_salary, 
    banks.name
	FROM wallets
    JOIN banks ON wallets.bank_id = banks.id
    JOIN permission_groups ON wallets.permission_group_id = permission_groups.id
    JOIN user_permission ON user_permission.permission_group_id = permission_groups.id
    JOIN users ON users.id = user_permission.user_id
    WHERE users.username = $1`

	rows, err := repo.db.QueryxContext(ctx, q, username)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			repo.log.Error("repo.wallets.GetByUsername", zap.Error(err))
		}
	}()

	var wallets = make([]models.WalletExtended, 0)

	for rows.Next() {
		var wallet models.WalletExtended

		err = rows.Scan(
			&wallet.ID,
			&wallet.Name,
			&wallet.Description,
			&wallet.Permission,
			&wallet.CreatedAt,
			&wallet.Currency,
			&wallet.IsSalary,
			&wallet.BankName,
		)

		if err != nil {
			return nil, err
		}

		wallets = append(wallets, wallet)
	}

	return wallets, err
}
