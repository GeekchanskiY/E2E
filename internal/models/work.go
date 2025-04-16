package models

import (
	"database/sql"
	"time"
)

type UserWork struct {
	ID         int64           `db:"id"`
	Name       string          `db:"name"`
	HourlyRate sql.NullFloat64 `db:"hourly_rate"`
	Worker     int64           `db:"worker"`
}

type WorkTime struct {
	ID        int64        `db:"id"`
	WorkID    int64        `db:"work_id"`
	StartTime time.Time    `db:"start_time"`
	EndTime   sql.NullTime `db:"end_time"`
}
