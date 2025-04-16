package models

type Bank struct {
	// Internal id for system
	ID int64 `db:"id" json:"id"`

	// Readable name
	Name string `db:"name" json:"name"`
}
