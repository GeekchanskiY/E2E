package models

import (
	"time"
)

type EventType string

const (
	RegistryLog   EventType = "log"
	RegistryEvent EventType = "event"
)

type Event struct {
	ID      int64     `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Event   EventType `db:"event" json:"event"`
	Content string    `db:"content" json:"content"`
	Time    time.Time `db:"time" json:"time"`
}
