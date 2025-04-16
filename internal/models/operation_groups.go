package models

import (
	"errors"
)

type OperationGroup struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	WalletID int64  `db:"wallet_id"`
}

func (o *OperationGroup) Validate() error {
	if o.Name == "" {
		return errors.New("name is required")
	}

	if o.WalletID == 0 {
		return errors.New("walletId is required")
	}

	return nil
}
