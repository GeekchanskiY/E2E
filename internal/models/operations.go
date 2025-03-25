package models

import (
	"errors"
	"time"
)

type Operation struct {
	Id               int64     `db:"id"`
	OperationGroupId int64     `db:"operation_group_id"`
	Time             time.Time `db:"time"`
	IsMonthly        bool      `db:"is_monthly"`
	IsConfirmed      bool      `db:"is_confirmed"`
	Amount           float64   `db:"amount"`
	InitiatorId      int64     `db:"initiator_id"`
}

func (op *Operation) Validate() error {
	if op.Time.IsZero() {
		return errors.New("time is empty")
	}

	if op.InitiatorId == 0 {
		return errors.New("initiator id is empty")
	}

	return nil
}
