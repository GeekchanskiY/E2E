package models

import (
	"database/sql"
	"time"
)

type UserWork struct {
	Id         int64           `db:"id"`
	Name       string          `db:"name"`
	HourlyRate sql.NullFloat64 `db:"hourly_rate"`
	Worker     int64           `db:"worker"`
}

type WorkTime struct {
	Id        int64        `db:"id"`
	WorkId    int64        `db:"work_id"`
	StartTime time.Time    `db:"start_time"`
	EndTime   sql.NullTime `db:"end_time"`
}
