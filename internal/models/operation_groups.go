package models

import (
	"errors"
)

type OperationGroup struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	WalletId int64  `db:"wallet_id"`
}

func (o *OperationGroup) Validate() error {
	if o.Name == "" {
		return errors.New("name is required")
	}

	if o.WalletId == 0 {
		return errors.New("walletId is required")
	}

	return nil
}
