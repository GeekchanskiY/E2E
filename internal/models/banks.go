package models

type Bank struct {
	// Internal id for system
	Id int `db:"id" json:"id"`

	// Readable name
	Name string `db:"name" json:"name"`
}
