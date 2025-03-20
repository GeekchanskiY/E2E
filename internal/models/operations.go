package models

import (
	"errors"
	"time"
)

type Operation struct {
	Id               int       `db:"id"`
	OperationGroupId int       `db:"operation_group_id"`
	IsConsumption    bool      `db:"is_consumption"`
	Time             time.Time `db:"time"`
	IsMonthly        bool      `db:"is_monthly"`
	IsConfirmed      bool      `db:"is_confirmed"`
	Amount           float64   `db:"amount"`
	InitiatorId      int       `db:"initiator_id"`
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
