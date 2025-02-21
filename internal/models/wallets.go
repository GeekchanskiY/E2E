package models

import "time"

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyBYN Currency = "BYN"
	CurrencyRUB Currency = "RUB"
)

type Wallet struct {
	Id                int       `db:"id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	PermissionGroupId int       `db:"permission_group_id"`
	CreatedAt         time.Time `db:"created_at"`
	Currency          Currency  `db:"currency"`
	IsSalary          bool      `db:"is_salary"`

	// BankId refers to internal bank id
	BankId int `db:"bank_id" json:"bank_id"`
}
