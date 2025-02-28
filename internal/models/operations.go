package models

import "time"

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
