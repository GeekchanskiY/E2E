package myfin

import (
	"time"
)

type Currency struct {
	BankName string    `json:"bank_name"`
	Name     string    `json:"name"`
	BuyUsd   float64   `json:"buy_usd"`
	SellUsd  float64   `json:"sell_usd"`
	Time     time.Time `json:"time"`
}
