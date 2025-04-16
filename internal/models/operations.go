package models

import (
	"errors"
	"time"
)

type Operation struct {
	ID               int64     `db:"id"`
	OperationGroupID int64     `db:"operation_group_id"`
	Time             time.Time `db:"time"`
	IsMonthly        bool      `db:"is_monthly"`
	IsConfirmed      bool      `db:"is_confirmed"`
	Amount           float64   `db:"amount"`
	InitiatorID      int64     `db:"initiator_id"`
}

func (op *Operation) Validate() error {
	if op.Time.IsZero() {
		return errors.New("time is empty")
	}

	if op.InitiatorID == 0 {
		return errors.New("initiator id is empty")
	}

	return nil
}

type OperationExtended struct {
	ID                 int64     `db:"id"`
	OperationGroupID   int64     `db:"operation_group_id"`
	OperationGroupName string    `db:"operation_group_name"`
	Time               time.Time `db:"time"`
	IsMonthly          bool      `db:"is_monthly"`
	IsConfirmed        bool      `db:"is_confirmed"`
	Amount             float64   `db:"amount"`
	InitiatorID        int64     `db:"initiator_id"`
	InitiatorName      string    `db:"initiator_name"`
}
