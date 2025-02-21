package models

type OperationGroup struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	WalletId int    `db:"wallet_id"`
}
