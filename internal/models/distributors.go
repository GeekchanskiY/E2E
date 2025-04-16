package models

import (
	"errors"
)

type Distributor struct {
	ID             int64   `db:"id"`
	Name           string  `db:"name"`
	SourceWalletID int64   `db:"source_wallet_id"`
	TargetWalletID int64   `db:"target_wallet_id"`
	Percent        float64 `db:"percent"`
}

func (d *Distributor) Validate() error {
	if d.Name == "" {
		return errors.New("name is required")
	}

	if d.SourceWalletID == 0 {
		return errors.New("source_wallet_id is required")
	}

	if d.TargetWalletID == 0 {
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
	ID               int64   `db:"id"`
	Name             string  `db:"name"`
	SourceWalletID   int64   `db:"source_wallet_id"`
	SourceWalletName string  `db:"source_wallet_name"`
	TargetWalletID   int64   `db:"target_wallet_id"`
	TargetWalletName string  `db:"target_wallet_name"`
	Percent          float64 `db:"percent"`
}
