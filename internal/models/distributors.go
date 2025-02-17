package models

type Distributor struct {
	Id             int     `db:"id"`
	Name           string  `db:"name"`
	SourceWalletId int     `db:"source_wallet"`
	TargetWalletId int     `db:"target_wallet"`
	Percent        float64 `db:"percent"`
}
