package models

// TODO: set id's to int64
type OperationGroup struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	WalletId int    `db:"wallet_id"`
}
