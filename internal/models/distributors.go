package models

type Distributor struct {
	Id             int     `db:"id"`
	Name           string  `db:"name"`
	SourceWalletId int     `db:"source_wallet_id"`
	TargetWalletId int     `db:"target_wallet_id"`
	Percent        float64 `db:"percent"`
}
