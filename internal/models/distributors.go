package models

import (
	"errors"
)

type Distributor struct {
	Id             int     `db:"id"`
	Name           string  `db:"name"`
	SourceWalletId int     `db:"source_wallet_id"`
	TargetWalletId int     `db:"target_wallet_id"`
	Percent        float64 `db:"percent"`
}

func (d *Distributor) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}

	if d.SourceWalletId == 0 {
		return errors.New("source_wallet_id is required")
	}

	if d.TargetWalletId == 0 {
		return errors.New("target_wallet_id is required")
	}

	if d.Percent <= 0 {
		return errors.New("percent is required and must be greater than zero")
	}

	if d.Percent > 100 {
		return errors.New("percent must be less than 100")
	}

	return nil
}

type DistributorExtended struct {
	Id               int     `db:"id"`
	Name             string  `db:"name"`
	SourceWalletId   int     `db:"source_wallet_id"`
	SourceWalletName string  `db:"source_wallet_name"`
	TargetWalletId   int     `db:"target_wallet_id"`
	TargetWalletName string  `db:"target_wallet_name"`
	Percent          float64 `db:"percent"`
}
