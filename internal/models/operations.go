package models

import "time"

type Operation struct {
	Id               int       `db:"id"`
	OperationGroupId int       `db:"operation_group"`
	IsConsumption    bool      `db:"is_consumption"`
	Time             time.Time `db:"time"`
	Amount           float64   `db:"amount"`
	Initiator        int       `db:"initiator"`
}
