package models

import (
	"time"
)

type CurrencyState struct {
	ID           int64     `db:"id" json:"id"`
	BankID       int64     `db:"bank_id" json:"bank_id"`
	CurrencyName string    `db:"currency_name" json:"currency_name"`
	SourceName   string    `db:"source_name" json:"source_name"`
	SellUsd      float64   `db:"sell_usd" json:"sell_usd"`
	BuyUsd       float64   `db:"buy_usd" json:"buy_usd"`
	Time         time.Time `db:"time" json:"time"`
}
