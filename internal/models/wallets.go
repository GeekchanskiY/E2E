package models

import (
	"errors"
	"fmt"
	"time"
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyBYN Currency = "BYN"
	CurrencyRUB Currency = "RUB"
)

type Wallet struct {
	Id                int64     `db:"id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	PermissionGroupId int64     `db:"permission_group_id"`
	CreatedAt         time.Time `db:"created_at"`
	Currency          Currency  `db:"currency"`
	IsSalary          bool      `db:"is_salary"`

	// BankId refers to internal bank id
	BankId int64 `db:"bank_id" json:"bank_id"`
}

type WalletExtended struct {
	Id          int64
	Name        string
	Description string
	Permission  string
	CreatedAt   time.Time
	Currency    Currency
	IsSalary    bool

	// BankId refers to internal bank id
	BankName string
}

func (w *WalletExtended) Validate() error {
	if w.Name == "" {
		return errors.New(fmt.Sprintf("name is required"))
	}

	if w.Currency == "" {
		return errors.New(fmt.Sprintf("currency is required"))
	}

	if w.Currency != CurrencyUSD && w.Currency != CurrencyEUR && w.Currency != CurrencyBYN && w.Currency != CurrencyRUB {
		return errors.New(fmt.Sprintf("currency is invalid"))
	}

	if w.Permission == "" {
		return errors.New(fmt.Sprintf("permission is required"))
	}

	return nil
}
